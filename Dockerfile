FROM golang:1.22.3-alpine3.19 AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./appbin

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /build/appbin ./
CMD ["./appbin"]