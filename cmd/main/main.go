package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"

	_ "github.com/dianerieck/GoogleCloudRun/docs"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/dianerieck/GoogleCloudRun/internal/service"
	"github.com/dianerieck/GoogleCloudRun/internal/utils"
	"github.com/go-chi/chi/v5"
)

// Pre-compilando a regex para melhor performance
var cepRegex = regexp.MustCompile(`^\d{8}$`)

func loadEnv() {
	for _, path := range []string{".env", "../.env", "../../.env"} {
		_ = godotenv.Load(path)
	}
}

func main() {

	// Carrega o .env se existir (útil localmente), mas não falha se estiver ausente (Cloud Run)
	loadEnv()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is not set. Configure it in Cloud Run or local .env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("starting server on port %s", port)

	r := chi.NewRouter()

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("/docs/swagger.json")))

	r.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No Dockerfile, a pasta docs é copiada para a raiz /docs
		http.ServeFile(w, r, "/docs/swagger.json")
	})

	r.Get("/cep/{cep}", func(w http.ResponseWriter, r *http.Request) {
		cepParam := chi.URLParam(r, "cep")

		if !cepRegex.MatchString(cepParam) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "invalid zipcode",
			})

			return
		}
		cidade, err := service.GetCidade(cepParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			// Verifica se o erro é de timeout de rede de forma mais robusta
			if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
				w.WriteHeader(http.StatusGatewayTimeout)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Timeout de 5 segundos excedido",
				})
				return
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "can not find zipcode",
			})
			return
		}

		celsius, err := service.GetTemperatura(cidade, apiKey)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
				w.WriteHeader(http.StatusGatewayTimeout)
				json.NewEncoder(w).Encode(map[string]string{
					"message": "Timeout de 5 segundos excedido - nenhuma API respondeu a tempo",
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Temp_C float64 `json:"temp_C"`
			Temp_F float64 `json:"temp_F"`
			Temp_K float64 `json:"temp_K"`
		}{
			Temp_C: celsius,
			Temp_F: utils.CelsiusToFahrenheit(celsius),
			Temp_K: utils.CelsiusToKelvin(celsius),
		})
	})

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
