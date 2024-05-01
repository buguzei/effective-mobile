FROM golang:1.22.1
LABEL authors="buguzei"

#WORKDIR /app

COPY ./ ./

COPY ../go.mod .

RUN go build -o ./bin/auth ./cmd/main.go

CMD ["./bin/auth"]