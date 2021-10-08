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


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /app/main /app/main
COPY --from=builder /app/etc /app/etc

WORKDIR /app
CMD ["./main", "-f", "etc/user-api.yaml"]
