FROM golang:1.17.2-alpine

WORKDIR /go/src/app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod tidy
EXPOSE 8080

CMD ["go", "run", "main.go"]