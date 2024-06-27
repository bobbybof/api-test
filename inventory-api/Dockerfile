FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GIN_MODE=release
ENV GO111MODULE=on

RUN go build -o main .

CMD ["./main"]