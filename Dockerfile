FROM golang:1.19.2-alpine3.15 AS builder 

WORKDIR /app
COPY  . .
RUN go build -o  main main.go 

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]