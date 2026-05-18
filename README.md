## Sistema de Clima por CEP - Go + Cloud Run
Aplicação desenvolvida em Golang que recebe um CEP válido de 8 dígitos, identifica a cidade correspondente e retorna o clima atual (temperatura em Celsius, Fahrenheit e Kelvin).
O sistema está publicado e acessível no Google Cloud Run.

---

## Requisitos Funcionais
- Entrada: CEP válido de 8 dígitos.
- Localização: Busca do CEP via ViaCEP.
- Clima: Consulta da temperatura atual via WeatherAPI.
- Conversão:
- Fahrenheit = Celsius × 1.8 + 32
- Kelvin = Celsius + 273

## Tecnologias utilizadas

- Go (Golang)
- Docker
- REST API
- Swagger (documentação)
- Google Cloud Run
- Testes automatizados (Go testing)


## Especificações da API
- Cenário 1: Sucesso
    HTTP 200
    json { "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }

- Cenário 2: Falha (Formato Inválido)
    HTTP 422
    json { "message": "invalid zipcode" }

- Cenário 3: Falha (CEP não encontrado)
    HTTP 404
    json { "message": "can not find zipcode" }


## Executando Localmente com Docker
Bash
   # Clonar repositório
        git clone https://github.com/dianerieck/GoogleCloudRun
        cd GoogleCloudRun

    # Build da imagem
        docker build -t cep-clima .

    # Rodar container
        docker run -p 8080:8080 cep-clima

## Endpoint Local
    http://localhost:8080/cep/{CEP}

## Deploy no Google Cloud Run
    A aplicação está disponível em: https://google-cloud-run-o5cyk2bmrq-uc.a.run.app/cep/99700054  

# Autora
Diane Rieck

# Licença

Projeto para fins educacionais e estudo.







