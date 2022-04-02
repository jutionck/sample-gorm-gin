FROM golang:alpine as build-env

RUN apk update && apk add git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o sample-go

FROM alpine
WORKDIR /app
COPY --from=build-env /src/sample-go /app

ENTRYPOINT ./sample-go