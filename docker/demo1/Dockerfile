FROM  golang:1.16-alpine AS go
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
RUN dep ensure -vendor-only

COPY . /go/src/project/

RUN go build -o /bin/project

FROM scratch

COPY --from=go /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]
