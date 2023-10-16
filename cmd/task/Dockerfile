FROM golang:1.21.2-alpine

RUN mkdir /app
WORKDIR /app

COPY ./cmd/task ./cmd/task
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

WORKDIR /app/cmd/task
RUN CGO_ENABLED=0 GOOS=linux go build -o task

CMD ["/app/cmd/task/task"]
