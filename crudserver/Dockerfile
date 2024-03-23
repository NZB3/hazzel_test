FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o crudserver cmd/crudserver/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/crudserver .
EXPOSE ${GOODS_API_PORT}
CMD ["./crudserver"]