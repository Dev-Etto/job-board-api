# Job Board API

## Visão Geral

A Job Board API é um serviço RESTful projetado para gerenciar vagas de emprego. Ela fornece endpoints para criar, ler, atualizar e deletar vagas de emprego, além de listar todas as vagas disponíveis. A API é construída em [Go](https://go.dev/) e utiliza o framework [Gin](https://gin-gonic.com/) para roteamento e [GORM](https://gorm.io/) para interações com o banco de dados.

## Funcionalidades

- **Create Job Openings**: Adicionar novas vagas de emprego com detalhes como cargo, empresa, localização, status remoto, link e salário.
- **List Job Openings**: Recuperar uma lista de todas as vagas de emprego.
- **Find Job Openings**: Obter detalhes de uma vaga específica pelo ID.
- **Update Job Openings**: Modificar detalhes de uma vaga existente.
- **Delete Job Openings**: Remover uma vaga pelo ID.
- **Swagger Documentation**: Documentação interativa da API disponível em `/swagger/index.html`.

## Estrutura do Projeto

```
job-board-api/
├── config/         # Arquivos de configuração (ex.: banco de dados, logger)
├── db/             # Arquivo do banco de dados SQLite
├── docs/           # Arquivos de documentação Swagger
├── handler/        # Handlers para os endpoints da API
├── router/         # Configuração de rotas da API
├── schema/         # Definições de esquema do banco de dados
├── main.go         # Ponto de entrada da aplicação
├── makefile        # Comandos para build e execução
├── go.mod          # Dependências do módulo Go
├── LICENSE         # Arquivo de licença
```

## Instalação

1. Clone o repositório:
   ```bash
   git clone <repository-url>
   cd job-board-api
   ```

## Uso

### Início Rápido

Para iniciar rapidamente o servidor da API com a documentação Swagger, utilize o comando:

```bash
make
```

### Comandos Disponíveis no Makefile

O arquivo `makefile` contém os seguintes comandos para facilitar o uso e a manutenção do projeto:

- **default**: Executa o comando `run-with-docs`, que inicializa a documentação Swagger e executa o servidor da API.

  ```bash
  make default
  ```

- **run**: Executa o servidor da API sem inicializar a documentação Swagger.

  ```bash
  make run
  ```

- **run-with-docs**: Inicializa a documentação Swagger, organiza as dependências do módulo Go e executa o servidor da API.

  ```bash
  make run-with-docs
  ```

- **build**: Compila o código-fonte e gera um binário executável chamado `job-board-api`.

  ```bash
  make build
  ```

- **test**: Executa os testes automatizados do projeto com saída detalhada.

  ```bash
  make test
  ```

- **docs**: Gera a documentação Swagger para a API.

  ```bash
  make docs
  ```

- **clean**: Remove o binário gerado e os arquivos de documentação.
  ```bash
  make clean
  ```

## Endpoints da API

### URL Base

`/api/v1`

### Endpoints

- **POST /opening**: Criar uma nova vaga de emprego.
- **GET /opening**: Listar todas as vagas de emprego.
- **GET /opening/{id}**: Buscar uma vaga de emprego pelo ID.
- **PUT /opening/{id}**: Atualizar uma vaga de emprego pelo ID.
- **DELETE /opening/{id}**: Deletar uma vaga de emprego pelo ID.

## Banco de Dados

A API utiliza [SQLite](https://www.sqlite.org/) como banco de dados. O arquivo do banco está localizado em `db/main.db`. O esquema é automaticamente migrado utilizando o [GORM](https://gorm.io/).

## Funcionalidades Futuras

Estamos comprometidos em expandir e melhorar continuamente a Job Board API. Aqui estão algumas funcionalidades que serão implementadas futuramente:

### Testes Automatizados

- **Cobertura de Testes**: Implementação de testes unitários e de integração para garantir a qualidade e a confiabilidade do código.
- **Framework de Testes**: Utilização de frameworks como [Testify](https://github.com/stretchr/testify) para facilitar a escrita e execução de testes.

### Novas Rotas

- **Candidatura do Candidato**: Adicionar endpoints para permitir que candidatos se candidatem às vagas disponíveis. Isso incluirá:

  - **POST /application**: Criar uma nova candidatura para uma vaga específica.
  - **GET /application**: Listar todas as candidaturas realizadas.
  - **GET /application/{id}**: Obter detalhes de uma candidatura específica pelo ID.
  - **DELETE /application/{id}**: Remover uma candidatura pelo ID.

- **Aprovação do Candidato**: Adicionar endpoints para gerenciar o status de aprovação dos candidatos. Isso incluirá:
  - **PUT /application/{id}/approve**: Aprovar um candidato para uma vaga específica.
  - **PUT /application/{id}/reject**: Rejeitar um candidato para uma vaga específica.

Essas melhorias visam tornar a API mais robusta e atender a um conjunto mais amplo de casos de uso, proporcionando uma experiência mais completa para os usuários.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
