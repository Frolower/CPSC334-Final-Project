FROM golang:1.24 AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .

RUN cd cmd && go build -o backend

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/cmd/backend .

EXPOSE 8080

CMD ["./backend"]
