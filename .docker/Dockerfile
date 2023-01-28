FROM golang:1.19

RUN mkdir /usr/src/app
WORKDIR /usr/src

RUN go install github.com/cosmtrek/air@v1.40.4

COPY ./go.mod ./go.sum ./.air.toml ./
RUN go mod download

COPY ./.docker/Dockerfile ./
COPY ./src ./

CMD ["air", "-c", ".air.toml"]
