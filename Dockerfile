FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app/movies

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/movies .

EXPOSE 8080

CMD ["./out/movies"]
