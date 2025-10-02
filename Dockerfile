# Use a imagem oficial do Go
FROM golang:1.25-alpine AS builder

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum primeiro para cache de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Instalar air para hot reload (versão compatível com Go 1.21)
RUN go install github.com/cosmtrek/air@v1.40.4

# Copiar o código fonte
COPY . .

# Expor a porta
EXPOSE 81

# Comando para executar com hot reload
CMD ["air", "-c", ".air.toml"]
