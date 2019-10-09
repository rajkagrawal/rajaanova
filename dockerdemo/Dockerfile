FROM golang:1.13.1-alpine3.10

WORKDIR /dockerdemo

COPY . .

RUN go build -o dockerdemo .
EXPOSE 8081

CMD ["./dockerdemo"]



