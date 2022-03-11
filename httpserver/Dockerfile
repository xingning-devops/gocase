FROM golang:1.17 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /httpserver
COPY . .
RUN go build -o httpserver .

FROM scratch
COPY --from=builder /httpserver/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]
