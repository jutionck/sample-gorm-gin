FROM golang:alpine

WORKDIR /app

COPY . .

RUN go build -o sample-go

CMD ./sample-go