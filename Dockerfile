FROM golang:dlv

COPY . /go/src/github.com/russellwmy/gin-boilerplate
WORKDIR /go/src/github.com/russellwmy/gin-boilerplate

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2"]