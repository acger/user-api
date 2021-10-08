FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/main ./main.go


FROM nginx:alpine

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/etc /app/etc
COPY  /nginx /etc/nginx/conf.d

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata && nginx -g daemon=on
ENV TZ Asia/Shanghai

CMD ["./main", "-f", "etc/user-api.yaml"]
