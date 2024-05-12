FROM golang:latest

WORKDIR /go/src/app

COPY ../go.mod .

COPY . .

RUN go build -o ./bin/cars ./cars/cmd/main.go
#RUN go install github.com/pressly/goose/v3/cmd/goose@latest
#RUN cd ./cars/migrations && goose postgres 'host=localhost port=5455 user=buguzei password=password dbname=cars sslmode=disable' up

CMD ["./bin/cars"]