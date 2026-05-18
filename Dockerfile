FROM golang:1.23-alpine AS builder

# Instala certificados CA (necessários para chamadas HTTPS para ViaCEP/WeatherAPI)
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY . .

# Compila apontando para o diretório correto onde está o main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cloudrun /cloudrun
COPY --from=builder /app/docs /docs
ENTRYPOINT ["./cloudrun"]