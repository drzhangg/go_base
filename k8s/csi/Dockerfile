FROM alpine

WORKDIR /build

ENV GO111MODULE=on \
	GOPROXY=https://goproxy.cn

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -a -o demo main.go

ENTRYPOINT ["/fsd-csi-driver"]