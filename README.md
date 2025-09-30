# Financial Track - Backend

> API REST para gerenciamento de gastos mensais pessoais. Permite cadastrar usuários, registrar despesas e acompanhar seu histórico financeiro.

---

## Tecnologias

- [Go](https://golang.org/)  
- [Gin](https://github.com/gin-gonic/gin) - framework web  
- [GORM](https://gorm.io/) - ORM para Go  
- [PostgreSQL](https://www.postgresql.org/) - banco de dados  
- [Docker](https://www.docker.com/) - para containerização do banco  
- [dotenv](https://github.com/joho/godotenv) - variáveis de ambiente  

---

## Estrutura do Projeto

financial-track-back/
├── cmd/
│   └── app.go                 # Arquivo principal para iniciar o servidor
├── controller/
│   └── user_controller.go     # Controlador para gerenciar ações de usuários
├── database/
│   └── main_database.go       # Configurações de conexão com o banco de dados
├── model/
│   └── user.go                # Definição do modelo de usuário
├── repository/
│   └── user_repository.go     # Repositório para interagir com o banco de dados de usuários
├── route/
│   ├── health.go              # Rota para verificar a saúde da API
│   └── user_routes.go         # Rota para endpoints relacionados a usuários
├── usecase/
│   └── user_usecase.go        # Lógica de negócios para usuários
├── utils/
│   └── validator.go           # Funções auxiliares de validação
├── middleware/
│   └── auth_middleware.go     # Middleware de autenticação JWT
├── migrations/
│   ├── 001_create_users.sql
│   └── 002_create_expenses.sql
├── .env.example               # Exemplo do arquivo .env
├── .env                       # Arquivo de variáveis de ambiente
├── .gitignore                 # Arquivos e pastas ignoradas pelo Git
├── docker-compose.yml         # Configuração do Docker para o banco de dados
├── go.mod                     # Gerenciamento de dependências do Go
├── go.sum                     # Checksums das dependências
└── README.md                  # Este arquivo

---

## Configuração do Ambiente

1. Clone este repositório:

   ```bash
   git clone https://github.com/erichenriquesm/financial-track-back.git
   cd financial-track-back

2. Copie o arquivo .env.example para .env e configure suas variáveis de ambiente, como as credenciais do banco de dados:

   ```bash
   cp .env.example .env

3. Instale as dependências do Go:

   ```bash
   go mod tidy

4. Inicie o banco de dados com Docker (se necessário):

   ```bash
   docker-compose up -d

4. Rode o servidor:

   ```bash
   go run cmd/app.go
