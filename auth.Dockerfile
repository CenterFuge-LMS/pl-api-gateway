FROM golang:alpine AS builder

LABEL maintainer="Safwanseban"

WORKDIR /grpc-gateway

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api-server


FROM alpine:latest

COPY --from=builder /grpc-gateway/main .

COPY . .

EXPOSE 3000

CMD [ "./main" ]