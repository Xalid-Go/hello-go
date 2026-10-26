FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /hello

FROM alpine:3.21
RUN adduser -D -u 1000 app
USER app
COPY --from=builder /hello /hello
CMD ["/hello"]
