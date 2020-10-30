FROM golang:alpine AS builder
RUN apk update
RUN apk add --virtual build-dependencies build-base gcc wget git
RUN mkdir /go/src/morningo
ADD . /go/src/morningo
WORKDIR /go/src/morningo
RUN make deps
RUN make build

FROM alpine:3.9
RUN apk update
RUN apk add --virtual build-dependencies build-base bash gcc musl-dev wget git openssl go
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash

RUN wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
RUN export GOPATH=/go
RUN export GOBIN=$GOPATH/bin
RUN export PATH=$PATH:$GOPATH/bin

ENV GOPATH=/go
ENV PATH $PATH:/go/bin

COPY --from=builder /go/src/morningo/build /go/bin

ADD . /go/src/morningo
WORKDIR /go/src/morningo

EXPOSE 9000
#ENTRYPOINT /go/bin/morningo --port 9000
CMD ["/go/bin/morningo"]


