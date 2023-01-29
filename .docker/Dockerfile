FROM golang:1.19

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.40.4

COPY ./go.mod ./go.sum  ./
RUN go mod download

COPY ./ ./

CMD ["air", "-c", ".air.toml"]
