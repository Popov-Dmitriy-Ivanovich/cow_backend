FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o run server.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/run /build/run
COPY --from=builder /build/.env /build/.env
EXPOSE 8080
CMD ["./run"]