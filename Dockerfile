FROM golang:1.16.4-buster AS builder

WORKDIR /app

RUN apt update && apt install -y upx

ADD main.go go.mod /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /go/bin/ping

RUN upx --brute /go/bin/ping


FROM scratch

COPY --from=builder /go/bin/ping /ping

ENTRYPOINT ["/ping"]
CMD ["/ping"]
