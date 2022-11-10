FROM golang:latest

ENV GIN_MODE=release

WORKDIR /app

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD [ "/app/main" ]