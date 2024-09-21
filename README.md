## Desafio #02 - Multithreading

O desafio consiste em criar uma aplicação que utiliza o mecanismo de Multithreading do Golang para pesquisar um cep em duas API's diferentes e apresentar a resposta da origem mais eficiência (menor latência).

#### 🖥️ Detalhes Gerais:

A busca do CEP será realizada em duas API's diferentes, [Viacep](https://viacep.com.br/) e [BrasilAPI](https://brasilapi.com.br/):
- Apenas o resultado da API que retornar em menor tempo será considerada.
- O tempo máximo de espera na resposta (timeout) da API é de 1 segundo.
- Em caso de falha na requisição por timeout do servidor, o erro será devidamente identificado e apresentado via console.
- Caso ambas as API's falhem (por CEP inválido, por exemplo) um erro será apresentado via console.
- Após o resultado com sucesso, os detalhes do endereço serão apresentados via console juntamente com a URL da API buscada.

#### 🗂️ Estrutura do Projeto
    .
    ├── cmd                # Entrypoint da aplicação
    ├── config             # helpers para configuração da aplicação
    ├── pkg                # Pacotes reutilizáveis utilizados na aplicação
    └── README.md

#### 🧭 Parametrização
A aplicação servidor possui um arquivo de configuração `config/config.go` onde é possível definir os parâmetros de timeout e URL's das API's para busca das informações do endereço.

```
API_REQUEST_1_URL=https://viacep.com.br/ws/<<zip>>/json           # URL da primeira API de consulta de CEP
API_REQUEST_2_URL=https://brasilapi.com.br/api/cep/v1/<<zip>>     # URL da segunda API de consulta de CEP
REQUEST_TIMEOUT=1                                                 # Tempo máximo de espera na resposta da API em segundos
```

> 💡 Caso necessário alterar as URL's das API's, basta alterar os valores das variáveis de ambiente `API_REQUEST_1_URL` e `API_REQUEST_2_URL` no arquivo `config/config.go`. Note que para que a lógica permaneça funcional, o local do CEP deve ser substituído por `<<zip>>`.

#### 🚀 Execução:
Para executar a aplicação, use o comando abaixo:
```bash
$ go mod tidy                 # Install dependencies if needed
$ go run cmd/main.go <<cep>>  # Replace <<cep>> with the desired zip code
```
### 📝 Exemplo de Execução:
```bash
$ go run cmd/main.go 01311000
2024/09/21 14:27:45 INFO : "Faster Response from"=https://brasilapi.com.br/api/cep/v1/01311000
2024/09/21 14:27:45 INFO :: "Response Data:"="{'cep':'01311000','city':'São Paulo','neighborhood':'Bela Vista','service':'open-cep','state':'SP','street':'Avenida Paulista'}"
```

```bash
$ go run cmd/main.go 01311000
2024/09/21 14:28:57 INFO : "Faster Response from"=https://viacep.com.br/ws/01311000/json
2024/09/21 14:28:57 INFO :: "Response Data:"="{'bairro':'Bela Vista','cep':'01311-000','complemento':'até 609 - lado ímpar','ddd':'11','estado':'São Paulo','gia':'1004','ibge':'3550308','localidade':'São Paulo','logradouro':'Avenida Paulista','regiao':'Sudeste','siafi':'7107','uf':'SP','unidade':''}"
```

```bash
$ go run cmd/main.go 01311000
2024/09/21 14:38:12 ERROR Timeout making request
```

```bash
$ go run cmd/main.go 01311032
2024/09/21 14:44:26 ERROR Error making request msg="not found result for the zip code"
2024/09/21 14:44:26 INFO : "Faster Response from"=https://viacep.com.br/ws/01311032/json
2024/09/21 14:44:26 INFO :: "Response Data:"={'erro':'true'}
```

### 🧪 Forçando delay nas API's:
Para simular um cenário onde a API 1 é mais lenta que a API 2, é possível forçar um delay na resposta da API 1. Para isso, basta alterar o valor do terceiro parâmetro na chamada da função `makeRequest()` no arquivo `cmd/main.go:32` ou `cmd/main.go:33`.

```go
func main() {
	...
    go makeRequest(url1, ch1, 0) # Adicionando um valor (int) no terceiro parâmetro, é possível forçar um delay em segundos na resposta da API1 (Viacep)
    go makeRequest(url2, ch2, 0) # Adicionando um valor (int) no terceiro parâmetro, é possível forçar um delay em segundos na resposta da API2 (BrasilAPI)
	...
}
```

> 💡 Caso o valor do delay de ambas as chamadas seja maior que o valor de `REQUEST_TIMEOUT` (`.env`), a resposta da execução deverá sempre incorrer em erro (timeout).