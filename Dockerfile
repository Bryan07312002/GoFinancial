FROM golang:1.24.2-alpine3.21

WORKDIR /app

COPY . .

EXPOSE 8080

RUN apk add --no-cache sqlite sqlite-dev gcc musl-dev

CMD ["go", "run", "/app/cmd/server/main.go"]
