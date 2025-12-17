# -------- Build stage --------
FROM golang:1.22-alpine AS builder

# Instala protoc e git (necessário para go install)
RUN apk add --no-cache protobuf git

# Instala plugins do protoc via Go
# Cria um módulo temporário para instalar os plugins
RUN mkdir -p /tmp/protoc-plugins && \
	cd /tmp/protoc-plugins && \
	go mod init temp && \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1 && \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adiciona Go bin ao PATH
ENV PATH="${PATH}:/root/go/bin"

WORKDIR /app

# Copia go.mod e go.sum
COPY go.mod go.sum ./
RUN go mod download

# Gera código gRPC do proto
COPY proto/price.proto ./proto/
RUN mkdir -p proto/gen/price && \
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/price.proto

# Copia o resto do código
COPY . .

# Gera código GraphQL
RUN go run github.com/99designs/gqlgen generate

# Build da aplicação
RUN go build -o gateway ./cmd/gateway

# -------- Runtime stage --------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/gateway ./gateway

EXPOSE 8080

ENTRYPOINT ["./gateway"]

