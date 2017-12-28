# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang  as builder

ADD . /go/src/golang-api

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/golang-api
RUN glide install
RUN go build && go install

ENTRYPOINT /go/bin/golang-api

EXPOSE 8080
