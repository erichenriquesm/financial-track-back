# Financial Track - Backend

> API REST para gerenciamento de gastos mensais pessoais. Permite cadastrar usuários, registrar despesas e acompanhar o histórico financeiro.

---

## Tecnologias

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) - framework web
- [GORM](https://gorm.io/) - ORM para Go
- [PostgreSQL](https://www.postgresql.org/) - banco de dados
- [Docker](https://www.docker.com/) - para containerização do banco
- [JWT](https://github.com/golang-jwt/jwt) - autenticação e autorização
- [bcrypt](https://golang.org/x/crypto/bcrypt) - criptografia de senhas
- [UUID](https://github.com/google/uuid) - geração de identificadores únicos
- [dotenv](https://github.com/joho/godotenv) - variáveis de ambiente

---

## Estrutura do Projeto

```text
financial-track-back/
├── cmd/
│   └── app.go                 # Arquivo principal para iniciar o servidor
├── controller/
│   ├── expense_controller.go  # Controlador para gerenciar ações de despesas
│   └── user_controller.go     # Controlador para gerenciar ações de usuários
├── database/
│   └── main_database.go       # Configurações de conexão com o banco de dados
├── middleware/
│   └── auth_middleware.go     # Middleware de autenticação JWT
├── model/
│   ├── expense.go             # Definição do modelo de despesa
│   └── user.go                # Definição do modelo de usuário
├── repository/
│   ├── expense_repository.go  # Repositório para interagir com o banco de dados de despesas
│   └── user_repository.go     # Repositório para interagir com o banco de dados de usuários
├── route/
│   ├── expense.go             # Rotas para endpoints relacionados a despesas
│   ├── health.go              # Rota para verificar a saúde da API
│   └── user.go                # Rotas para endpoints relacionados a usuários
├── usecase/
│   ├── expense.go             # Lógica de negócios para despesas
│   └── user.go                # Lógica de negócios para usuários
├── utils/
│   ├── auth.go                # Funções auxiliares de autenticação
│   ├── JSONTime.go            # Utilitário para formatação de tempo JSON
│   └── validator.go           # Funções auxiliares de validação
├── .env.example               # Exemplo do arquivo .env
├── .env                       # Arquivo de variáveis de ambiente
├── .gitignore                 # Arquivos e pastas ignoradas pelo Git
├── docker-compose.yml         # Configuração do Docker para o banco de dados
├── go.mod                     # Gerenciamento de dependências do Go
├── go.sum                     # Checksums das dependências
└── README.md                  # Este arquivo

---

## Funcionalidades

### Autenticação
- **POST** `/auth/register` - Cadastro de usuários
- **POST** `/auth/login` - Login de usuários

### Despesas (Autenticação necessária)
- **POST** `/expenses/` - Criar nova despesa
- **GET** `/expenses/` - Listar despesas do usuário

### Health Check
- **GET** `/ping` - Verificar status da API

## Categorias de Despesas

O sistema suporta as seguintes categorias:
- `FOOD` - Alimentação
- `TRANSPORTATION` - Transporte
- `HOUSING` - Moradia
- `HEALTH` - Saúde
- `EDUCATION` - Educação
- `ENTERTAINMENT` - Entretenimento
- `CLOTHING` - Roupas
- `PERSONAL` - Pessoal
- `FINANCE` - Financeiro
- `OTHERS` - Outros

---

## Configuração do Ambiente

1. Clone este repositório:

   git clone https://github.com/erichenriquesm/financial-track-back.git
   cd financial-track-back

2. Copie o arquivo .env.example para .env e configure suas variáveis de ambiente:

   cp .env.example .env

   Configure as seguintes variáveis no arquivo `.env`:
   ```
   DB_URL=postgres://username:password@localhost:5432/financial_track?sslmode=disable
   JWT_SECRET=your-secret-key-here
   SERVER_PORT=81
   ```

3. Instale as dependências do Go:

   go mod tidy

4. Inicie o banco de dados com Docker:

   docker-compose up -d

5. Execute o servidor:

   go run cmd/app.go

   O servidor estará disponível em `http://localhost:81` (ou na porta configurada em `SERVER_PORT`)

## Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `DB_URL` | URL de conexão com o PostgreSQL | - |
| `JWT_SECRET` | Chave secreta para assinatura dos tokens JWT | - |
| `SERVER_PORT` | Porta do servidor | `81` |
