FROM golang:1.23-alpine AS builder

# Instala certificados CA (necessários para chamadas HTTPS para ViaCEP/WeatherAPI)
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Habilitado ANTES do download para evitar erros de conexão (unexpected EOF)
# Copia apenas os arquivos de dependências primeiro para aproveitar o cache de camadas do Docker.
# Isso evita baixar todas as bibliotecas novamente se você alterar apenas o código-fonte.
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Compila apontando para o diretório correto onde está o main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/cloudrun /cloudrun
COPY --from=builder /app/docs /docs
ENTRYPOINT ["/cloudrun"]
