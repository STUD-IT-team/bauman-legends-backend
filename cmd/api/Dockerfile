FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY ./cmd/api ./cmd/api
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

WORKDIR /app/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o api

CMD ["/app/cmd/api/api"]