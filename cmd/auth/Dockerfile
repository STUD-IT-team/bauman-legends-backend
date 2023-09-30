FROM golang:latest

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/cmd/auth
RUN CGO_ENABLED=0 GOOS=linux go build -o auth

CMD ["/app/cmd/auth/auth"]