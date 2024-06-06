FROM golang:1.21.0-alpine3.18 as builder
WORKDIR /app/
COPY . .
RUN go mod download
RUN go build -o ./binary/app ./cmd

FROM alpine:3.18
WORKDIR /app/
COPY --from=builder /app/ .
ENTRYPOINT ["/app/binary/app"]