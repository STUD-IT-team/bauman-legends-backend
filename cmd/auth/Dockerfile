FROM golang:1.21.2-alpine

RUN mkdir /app
WORKDIR /app

COPY ./cmd/auth ./cmd/auth
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

WORKDIR /app/cmd/auth
RUN CGO_ENABLED=0 GOOS=linux go build -o auth

CMD ["/app/cmd/auth/auth"]