package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/dianerieck/CURSO-GO/11-DESAFIO-CEP/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/dianerieck/CURSO-GO/11-DESAFIO-CEP/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/swagger.json"),
	))

	r.Get("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "docs/swagger.json")
	})

	r.Get("/cep/{cep}", func(w http.ResponseWriter, r *http.Request) {
		cepParam := chi.URLParam(r, "cep")
		if cepParam == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, source, latency := service.FetchCEPRace(cepParam)
		if data == nil {
			w.WriteHeader(http.StatusGatewayTimeout)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(struct {
			Source  string      `json:"source"`
			Latency string      `json:"latency"`
			Data    interface{} `json:"data"`
		}{
			Source:  source,
			Latency: latency.String(),
			Data:    data,
		})
		w.Write(b)
	})

	http.ListenAndServe(":8000", r)
}
