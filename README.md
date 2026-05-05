# Desafio CEP - Multithreading com Go

Projeto desenvolvido em Golang para buscar informações de endereço a partir de um CEP para consultar múltiplas APIs simultaneamente e retornar a resposta mais rápida.

Este projeto foi criado como solução para um desafio técnico envolvendo:

- Multithreading em Go
- Consumo de APIs externas
- Controle de timeout
- Race condition entre serviços

---

## Objetivo

Realizar requisições simultâneas para duas APIs públicas de CEP:

- https://brasilapi.com.br
- http://viacep.com.br

A aplicação:

- Executa chamadas em paralelo  
- Retorna a resposta da API mais rápida  
- Descarta a resposta mais lenta  
- Limita o tempo total de resposta em 1 segundo  


## Tecnologias utilizadas

- Go (Golang)
- HTTP Client
- REST API
- Swagger (documentação)


## Como executar o projeto e Clonar o repositório

```bash
git clone https://github.com/dianerieck/Multithreading-Golang
cd seu-repo

Rodar aplicação
go run cmd/main/main.go

 
Endpoint:
http://localhost:8000/cep/99700088


Swagger (Documentação da API)
http://localhost:8000/swagger/index.html


## Testes via terminal
Teste simples
curl http://localhost:8000/cep/99700088


# Conceitos aplicados

Concorrência em Go
Race entre serviços externos
Timeout controlado
Clean architecture básica
Separação de responsabilidades


# Autora
Diane Rieck

# Licença

Projeto para fins educacionais e estudo.







