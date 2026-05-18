package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/dianerieck/GoogleCloudRun/docs"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/dianerieck/GoogleCloudRun/internal/service"
	"github.com/dianerieck/GoogleCloudRun/internal/utils"
	"github.com/go-chi/chi/v5"
)

// Pre-compilando a regex para melhor performance
var cepRegex = regexp.MustCompile(`^\d{8}$`)

func main() {

	_ = godotenv.Load("../../.env")

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
		path := "./docs/swagger.json"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			path = filepath.Join("..", "..", "docs", "swagger.json")
		}
		http.ServeFile(w, r, path)
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
			if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline") {
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
			if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline") {
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
