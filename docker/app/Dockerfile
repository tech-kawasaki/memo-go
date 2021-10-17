FROM golang:1.17.2-alpine

WORKDIR /go/src/app

RUN apk update && apk add git
COPY go.mod go.sum ./
ARG CGO_ENABLED=0
RUN go mod tidy && \
    go install github.com/rubenv/sql-migrate/...@latest
EXPOSE 8080

CMD ["go", "run", "main.go"]