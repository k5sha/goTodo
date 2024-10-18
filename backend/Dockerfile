FROM golang:1.23-bookworm


WORKDIR /todo-backend


COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o / ./cmd/go-todo


ENV CONFIG_PATH=./config/development.yaml

EXPOSE 8080

CMD ["/go-todo"]