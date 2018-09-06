FROM golang:latest as builder
WORKDIR /go/src/github.com/yuki-toida/refodpt/backend/
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    CGO_ENABLED=0 GOOS=linux GIN_MODE=release ENV=prod go build -o app .

FROM alpine:latest
EXPOSE 8010
ENV ENV=prod
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata
WORKDIR /opt/app
COPY --from=builder /go/src/github.com/yuki-toida/refodpt/backend/app .
COPY --from=builder /go/src/github.com/yuki-toida/refodpt/backend/.env .
COPY --from=builder /go/src/github.com/yuki-toida/refodpt/backend/config.prod.toml .
CMD ["./app"]
