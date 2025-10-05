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
│   ├── expense.go             # Modelo de despesa e DTOs
│   ├── pagination.go          # Structs de paginação reutilizáveis
│   ├── time.go                # Tipo JSONTime (parse/serialize no fuso)
│   └── user.go                # Modelo de usuário
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
  - Body (JSON, camelCase):
    ```json
    {
      "userId": "<uuid>",
      "category": "FOOD",
      "amount": 105,
      "description": "Lanche",
      "transactionAt": "2025-10-05 17:19" // Formato 2006-01-02 15:04
    }
    ```
- **GET** `/expenses/mensal-summary` - Resumo/paginação dos últimos 30 dias
  - Query params: `page`, `perPage`
  - Resposta (Laravel-like):
    ```json
    {
      "amount": 123.45,
      "data": [ /* expenses */ ],
      "currentPage": 1,
      "lastPage": 5,
      "totalItems": 87,
      "perPage": 20
    }
    ```

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

2. Crie o arquivo `.env` e configure suas variáveis de ambiente (exemplo abaixo):

   **Para desenvolvimento local, altere no .env:**
   ```
   DB_URL=postgres://admin:secret@localhost:5432/financial_track?sslmode=disable
   ```

   **Exemplo de .env (Docker/local):**
   ```
   DB_URL=postgres://admin:secret@financial_track_db:5432/financial_track?sslmode=disable
   JWT_SECRET=your-secret-key-here
   SERVER_PORT=81
   # Timezone da aplicação e da sessão do banco (padrão America/Sao_Paulo)
   APP_TIMEZONE=America/Sao_Paulo
   ```

3. Instale as dependências do Go:

   go mod tidy

4. Execute com Docker (recomendado):

   # Subir todos os serviços (banco + API)
   docker-compose up -d

   # Para ver os logs em tempo real
   docker-compose up

   O servidor estará disponível em `http://localhost:81`

5. Ou execute localmente (desenvolvimento):

   # Apenas o banco de dados
   docker-compose up financial_track_db -d

   # Execute o servidor localmente
   go run cmd/app.go

   O servidor estará disponível em `http://localhost:81`

## Desenvolvimento com Docker

### Hot Reload
O projeto está configurado com **hot reload** usando Air. Qualquer alteração no código será automaticamente refletida sem precisar rebuildar o container.

**Configuração otimizada:**
- **Polling**: Habilitado para detectar mudanças em containers
- **Symlinks**: Seguidos para melhor detecção
- **Rerun**: Automático quando há mudanças
- **Interrupt**: Envia sinal para parar processo anterior

**Para testar o hot reload:**
```bash
# Ver logs da API
make logs-api

# Em outro terminal, modifique um arquivo .go
# O Air deve detectar e rebuildar automaticamente
```

### Comandos Úteis

#### **Usando Makefile:**
```bash
# Subir containers
make up

# Parar containers
make down

# Reiniciar containers
make restart

# Ver logs em tempo real
make logs

# Modo desenvolvimento
make dev

# Status dos containers
make status
```

#### **Comandos Docker Diretos:**
```bash
# Subir todos os serviços
docker compose up -d

# Ver logs em tempo real
docker compose up

# Parar todos os serviços
docker compose down

# Rebuildar apenas a API
docker compose build financial_track_api

# Executar comandos no container
docker compose exec financial_track_api sh

# Ver logs apenas da API
docker compose logs -f financial_track_api
```

### Makefile - Comandos Disponíveis

O projeto inclui um Makefile com comandos úteis para desenvolvimento:

```bash
# Ver todos os comandos disponíveis
make help

# Comandos principais
make up          # Subir containers
make down         # Parar containers
make restart      # Reiniciar containers
make build        # Rebuildar containers
make logs         # Ver logs em tempo real
make dev          # Modo desenvolvimento
make prod         # Modo produção

# Comandos de manutenção
make clean        # Limpar cache do Docker
make status       # Status dos containers
make exec-api     # Executar comando na API
make exec-db      # Executar comando no banco

# Comandos de desenvolvimento
make install      # Instalar dependências
make test         # Executar testes
make fmt          # Formatar código
make lint         # Executar lint
```

### Resolução de Problemas

Se encontrar erro de versão do Go:

```bash
# Limpar cache do Docker
make clean

# Rebuildar sem cache
make build-no-cache
```

### Estrutura dos Containers
- **financial_track_db**: PostgreSQL na porta 5432
- **financial_track_api**: API Go na porta 81

### Configuração do Docker
O Docker Compose está configurado para ler as variáveis do arquivo `.env`.

**DSN para Docker**: `postgres://admin:secret@financial_track_db:5432/financial_track?sslmode=disable`

**Variáveis necessárias no .env:**
- `DB_URL` - URL de conexão com o banco
- `JWT_SECRET` - Para autenticação
- `SERVER_PORT` - Para a API
- `APP_TIMEZONE` - Fuso horário da aplicação/banco (padrão `America/Sao_Paulo`)

## Variáveis de Ambiente

| Variável | Descrição | Valor Padrão |
|----------|-----------|--------------|
| `DB_URL` | URL de conexão com o PostgreSQL | `postgres://admin:secret@financial_track_db:5432/financial_track?sslmode=disable` |
| `JWT_SECRET` | Chave secreta para assinatura dos tokens JWT | `your-secret-key-here` |
| `SERVER_PORT` | Porta do servidor | `81` |
| `APP_TIMEZONE` | Fuso horário da aplicação/banco | `America/Sao_Paulo` |

### **Para Desenvolvimento Local:**
Altere apenas a `DB_URL` no arquivo `.env`:
```
DB_URL=postgres://admin:secret@localhost:5432/financial_track?sslmode=disable
```
