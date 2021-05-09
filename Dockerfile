FROM golang:1.15.12-alpine3.12

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build src/app/main.go

EXPOSE 8080
ENTRYPOINT ["./main"]
