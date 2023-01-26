FROM golang:1.19

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@v1.40.4

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main ./src/cmd/main.go

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
