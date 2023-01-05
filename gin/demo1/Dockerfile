FROM golang:1.17.3-alpine3.13 as builder

WORKDIR /build

ENV GO111MODULE=on \
	GOPROXY=https://goproxy.cn

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -o demo main.go

FROM alpine:3.11 as final

WORKDIR /app

COPY --from=builder /build/demo /app/

RUN chmod +x /app/demo

ENV TZ=Asia/Shanghai

EXPOSE 8181

CMD ["/app/demo"]
