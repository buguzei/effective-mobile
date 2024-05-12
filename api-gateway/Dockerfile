FROM golang:latest

WORKDIR /go/src/app

COPY ../go.mod .

COPY . .

RUN go build -o ./bin/auth ./auth/cmd/main.go

CMD ["./bin/auth"]