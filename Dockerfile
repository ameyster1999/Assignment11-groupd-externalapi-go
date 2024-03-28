FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Assignment11-groupd-externalapi-go .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/Assignment11-groupd-externalapi-go .

EXPOSE 8014

CMD ["./Assignment11-groupd-externalapi-go"]