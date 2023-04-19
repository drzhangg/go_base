FROM golang:1.17.5-alpine3.15 as builder

WORKDIR /build

ENV GO111MODULE=on \
	GOPROXY=https://goproxy.cn

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -o csi main.go

#RUN cp csi /app/csi

#RUN chmod +x csi
#
#ENTRYPOINT ["csi"]


FROM alpine:3.11 as final

WORKDIR /app

COPY --from=builder /build/csi /app/

RUN chmod +x /app/csi

#ENV TZ=Asia/Shanghai

CMD ["/app/csi"]