FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go get github.com/julienschmidt/httprouter
RUN go build ./src/cmd/main.go

EXPOSE 8080
CMD ["app"]
