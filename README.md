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

    # Crie um arquivo .env com sua API_KEY ou passe via linha de comando
    # API_KEY=sua_chave_aqui

    # Build da imagem
        docker build -t cep-clima .

    # Rodar container passando a chave de API como variável de ambiente
    # Substitua 'SUA_CHAVE' pela chave real da WeatherAPI
    docker run -p 8080:8080 -e API_KEY=SUA_CHAVE cep-clima

## Endpoint Local
    http://localhost:8080/cep/{CEP}

## Deploy no Google Cloud Run
    A aplicação está disponível em: 
    https://google-cloud-run-o5cyk2bmrq-uc.a.run.app/cep/99700054  

### Configuração no Cloud Run
Ao fazer o deploy, certifique-se de configurar a variável de ambiente:
1. Vá para o Console do Cloud Run.
2. Selecione seu serviço -> "Edit & Deploy New Revision".
3. Na aba "Variables & Secrets", adicione `API_KEY` com o valor da sua chave.

# Autora
Diane Rieck

# Licença

Projeto para fins educacionais e estudo.
