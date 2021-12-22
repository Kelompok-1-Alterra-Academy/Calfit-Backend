FROM golang:1.16-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main

FROM alpine:3.14
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
EXPOSE 8000
CMD [ "./main" ]