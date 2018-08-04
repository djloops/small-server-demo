FROM golang

COPY . /go/src/github.com/djloops/small-server-demo
WORKDIR /go/src/github.com/djloops/small-server-demo
#
#CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2"]

CMD go run main.go