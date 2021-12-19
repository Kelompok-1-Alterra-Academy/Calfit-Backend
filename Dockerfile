FROM golang:1.17.5-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go clean --modcache
RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]