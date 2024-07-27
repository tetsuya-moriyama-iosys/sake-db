# backend/.devcontainer/Dockerfile
FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN go build -o /backend

EXPOSE 8080

CMD ["/backend"]
