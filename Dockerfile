# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang  as builder

ADD . /go/src/github.com/yogeshsr/search-spike

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/yogeshsr/search-spike
RUN glide install
RUN go build && go install

ENTRYPOINT /go/bin/search-spike

EXPOSE 8080
