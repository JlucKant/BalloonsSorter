FROM golang:alpine AS builder
WORKDIR /project
COPY . .

WORKDIR /project/app
RUN go build -o /app/main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

ENTRYPOINT ["/app/main"]