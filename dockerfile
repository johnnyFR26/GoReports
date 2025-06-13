# Etapa 1 - Build
FROM golang:1.22 AS builder

WORKDIR /app

# Copia os arquivos go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o código restante
COPY . .

# Compila o binário
RUN go build -o app .

# Etapa 2 - Imagem final
FROM alpine:latest

WORKDIR /root/

# Instala as dependências necessárias (para sqlite e segurança)
RUN apk --no-cache add ca-certificates sqlite

# Copia o binário da etapa de build
COPY --from=builder /app/app .

# Copia o banco de dados (se desejar)
COPY database.db ./

# Expõe a porta
EXPOSE 8080

# Comando para rodar
CMD ["./app"]
