FROM golang:1.22-alpine AS builder
WORKDIR /usr/local/src

COPY ./ ./
RUN go mod download

RUN go build -o ./bin/main /usr/local/src/cmd/main.go

FROM alpine:latest AS runner

COPY --from=builder /usr/local/src/bin/main /main

CMD ["/main"]
