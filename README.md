# Backend conta corrent

Projeto criado para fins de aprendizado e evolução de conhecimentos

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

### 📋 Pré-requisitos

Ferramentas:

- [Docker](https://docs.docker.com/desktop/)
- [Golang](https://golang.org/doc/install)
- [Migration](https://github.com/golang-migrate/migrate)
- [Swagger](https://github.com/swaggo/echo-swagger)

## 📦 Desenvolvimento

Existe dois comandos básicos para execução do projeto:

- `make run-docker`: Faz o build do projeto e sobe no docker, junto com todas as suas repectivas dependencias.
- `make dev`: wrapper para o `go run main.go`.

Existe dois comandos básicos para testar integridade do projeto:

- `make test`: realiza o teste unitario de todos os metodos.
- `make test-cover`: realiza o mesmo que o `make teste`, porém gera um relatorio em hml para melhor visualização.

## 🗂 Arquitetura

### Descrição dos diretórios e arquivos mais importantes:

- `./api`: O codígo relacionado com as rotas, middlewares e versionamento da api.
- `./api/api.go`: Nesse arquivo está toda parte de registros dos sub-modulos que existem nesse diretório, incluindo versionamento de rotas e gerenciamento de middlewares.
- `./api/v1`: Este diretório possui a configuração e registro de todos os sub-modulos.
- `./api/v1/v1.go`: Nesse arquivo está toda parte de registros dos sub-modulos que existem nesse diretório com o path `/v1/**`.
- `./api/middleware`: Aqui é aonde se encontra os middlewares em geral, como exemplo podemos citar os de injeção de sessão no contexto e o de autorização das rotas.

- `./app`: Este diretório possui a configuração e registro de todos os sub-modulos. Aqui se encontra todo o código que é utilizado para orquestração e regras de negôcio do serviço.
- `./app/app.go`: Arquivo para o registro, configuração e injeção de depêndencias externas nos sub-modulos..

- `./db/migrations`: Esse diretório possui todas as migration que serão necessarias para rodar a aplicação.

- `./model`: Este diretório possui todos os arquivos de modelos globais do projeto

- `./store`: Aqui se encontra todo o código que é utilizado para consultas usando banco de dados.
- `./store/store.go`: Arquivo para o registro, configuração e injeção de depêndencias como banco de dados.

- `./util`: Sub-modulos necessários para manutenção do projeto em geral.
- `./docs`: Arquivos gerados pelo swagger, referente a documentação.

## 🛠️ Construído com

- [echo](https://echo.labstack.com/) - Framework Web
- [go mod](https://blog.golang.org/using-go-modules) - Dependência
- [goDotEnv](https://github.com/joho/godotenv) - Configuração
- [logrus](github.com/sirupsen/logrus) - Log
- [migration](https://github.com/golang-migrate/migrate) - crianção da estrutura inicial do db (migration)
- [postgresql](https://www.postgresql.org/docs/) - banco de dados
