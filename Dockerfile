FROM golang

WORKDIR /app

COPY . .

RUN go build -o sample-go

EXPOSE 8081

CMD ./sample-go