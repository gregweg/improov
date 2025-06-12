# backend/Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o improov-server main.go

EXPOSE 8080
CMD ["./improov-server"]