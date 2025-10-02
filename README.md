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

2. Copie o arquivo env.example para .env e configure suas variáveis de ambiente:

   cp env.example .env

   **Para desenvolvimento local, altere no .env:**
   ```
   DB_URL=postgres://admin:secret@localhost:5432/financial_track?sslmode=disable
   ```

   **Para Docker, as variáveis já estão configuradas no env.example:**
   ```
   DB_URL=postgres://admin:secret@financial_track_db:5432/financial_track?sslmode=disable
   JWT_SECRET=your-secret-key-here
   SERVER_PORT=81
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

## Variáveis de Ambiente

| Variável | Descrição | Local | Docker |
|----------|-----------|-------|--------|
| `DB_URL` | URL de conexão com o PostgreSQL | `localhost:5432` | `financial_track_db:5432` |
| `DB_HOST` | Host do banco de dados | `localhost` | `financial_track_db` |
| `DB_PORT` | Porta do banco de dados | `5432` | `5432` |
| `DB_USER` | Usuário do banco de dados | `admin` | `admin` |
| `DB_PASSWORD` | Senha do banco de dados | `secret` | `secret` |
| `DB_NAME` | Nome do banco de dados | `financial_track` | `financial_track` |
| `JWT_SECRET` | Chave secreta para assinatura dos tokens JWT | - | - |
| `SERVER_PORT` | Porta do servidor | `81` | `81` |
