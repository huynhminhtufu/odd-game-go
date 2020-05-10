FROM golang:1.13-alpine3.10 as builder
RUN apk add --no-cache --update bash git
WORKDIR /src
ADD ./go.mod ./go.sum ./
RUN go mod download
ADD ./ ./
RUN go build -o /dist/dd-game-server cmd/server/*.go

FROM alpine:3.9
RUN apk add --no-cache --update ca-certificates tzdata curl && rm -rf /var/cache/apk/*
RUN cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
RUN echo "Asia/Ho_Chi_Minh" >  /etc/timezone

COPY --from=builder /dist/odd-game-server /app/bin/odd-game-server
WORKDIR /app/bin
CMD ["/app/bin/odd-game-server"]
