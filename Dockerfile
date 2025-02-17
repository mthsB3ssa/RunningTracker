FROM golang:1.21

WORKDIR /app

# Instala o CompileDaemon para reinicialização automática
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copia os arquivos de dependências e baixa os módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

EXPOSE 8080

# O CompileDaemon monitora mudanças, recompila e executa a aplicação
CMD CompileDaemon -build="go build -buildvcs=false -o main ./cmd/api" -command=./main


