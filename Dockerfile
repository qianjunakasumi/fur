FROM golang:1.19-alpine as builder
WORKDIR /usr/src/fur
COPY . .
RUN go build -ldflags "-s -w" -o ./fur ./...

FROM alpine:3
COPY --from=builder /usr/src/fur/fur /usr/local/bin/fur
CMD ["fur"]
