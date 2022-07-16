FROM golang:alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

WORKDIR /usr/src/app/server

RUN go build ./...

CMD ["./server"]




