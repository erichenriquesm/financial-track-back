# Financial Track - Docker Management
# Comandos para gerenciar containers

.PHONY: help up down restart build logs clean dev prod

# Comando padrão
help:
	@echo "Financial Track - Docker Management"
	@echo ""
	@echo "Comandos disponíveis:"
	@echo "  make up      - Subir containers em background"
	@echo "  make down    - Parar e remover containers"
	@echo "  make restart - Reiniciar containers"
	@echo "  make build   - Rebuildar containers"
	@echo "  make logs    - Ver logs em tempo real"
	@echo "  make clean   - Limpar cache do Docker"
	@echo "  make dev     - Modo desenvolvimento (logs visíveis)"
	@echo "  make prod    - Modo produção (background)"
	@echo ""

# Subir containers em background
up:
	@echo "🚀 Subindo containers..."
	docker compose up -d
	@echo "✅ Containers iniciados!"
	@echo "📊 API disponível em: http://localhost:81"
	@echo "🗄️  Banco disponível em: localhost:5432"

# Parar e remover containers
down:
	@echo "🛑 Parando containers..."
	docker compose down
	@echo "✅ Containers parados!"

# Reiniciar containers
restart:
	@echo "🔄 Reiniciando containers..."
	docker compose restart
	@echo "✅ Containers reiniciados!"

# Rebuildar containers
build:
	@echo "🔨 Rebuildando containers..."
	docker compose build
	@echo "✅ Containers rebuildados!"

# Rebuildar sem cache
build-no-cache:
	@echo "🔨 Rebuildando containers (sem cache)..."
	docker compose build --no-cache
	@echo "✅ Containers rebuildados!"

# Ver logs em tempo real
logs:
	@echo "📋 Exibindo logs em tempo real..."
	docker compose logs -f

# Ver logs apenas da API
logs-api:
	@echo "📋 Exibindo logs da API..."
	docker compose logs -f financial_track_api

# Ver logs apenas do banco
logs-db:
	@echo "📋 Exibindo logs do banco..."
	docker compose logs -f financial_track_db

# Limpar cache do Docker
clean:
	@echo "🧹 Limpando cache do Docker..."
	docker system prune -f
	@echo "✅ Cache limpo!"

# Modo desenvolvimento (logs visíveis)
dev:
	@echo "🔧 Iniciando modo desenvolvimento..."
	docker compose up

# Modo produção (background)
prod:
	@echo "🚀 Iniciando modo produção..."
	docker compose up -d
	@echo "✅ Aplicação rodando em background!"

# Status dos containers
status:
	@echo "📊 Status dos containers:"
	docker compose ps

# Executar comando no container da API
exec-api:
	@echo "🔧 Executando comando no container da API..."
	docker compose exec financial_track_api sh

# Executar comando no container do banco
exec-db:
	@echo "🔧 Executando comando no container do banco..."
	docker compose exec financial_track_db psql -U admin -d financial_track

# Testar hot reload
test-reload:
	@echo "🔄 Testando hot reload..."
	@echo "Modifique um arquivo .go e veja se o Air detecta a mudança"
	@echo "Use 'make logs-api' para ver os logs do Air"

# Backup do banco
backup:
	@echo "💾 Fazendo backup do banco..."
	docker compose exec financial_track_db pg_dump -U admin financial_track > backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "✅ Backup criado!"

# Restaurar backup do banco
restore:
	@echo "📥 Restaurando backup do banco..."
	@read -p "Digite o nome do arquivo de backup: " file; \
	docker compose exec -T financial_track_db psql -U admin -d financial_track < $$file
	@echo "✅ Backup restaurado!"

# Instalar dependências
install:
	@echo "📦 Instalando dependências..."
	go mod tidy
	@echo "✅ Dependências instaladas!"

# Executar testes
test:
	@echo "🧪 Executando testes..."
	go test ./...
	@echo "✅ Testes concluídos!"

# Formatar código
fmt:
	@echo "🎨 Formatando código..."
	go fmt ./...
	@echo "✅ Código formatado!"

# Lint do código
lint:
	@echo "🔍 Executando lint..."
	golangci-lint run
	@echo "✅ Lint concluído!"
