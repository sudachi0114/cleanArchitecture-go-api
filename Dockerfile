FROM golang:1.15.12-alpine3.12

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8080

RUN go get -u github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]
