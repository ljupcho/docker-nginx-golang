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
RUN apk add --virtual build-dependencies build-base gcc wget git
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
RUN adduser -S -D -H -h /go/bin appuser
USER appuser
COPY --from=builder /go/src/morningo/build /go/bin
ADD . /go/src/morningo
WORKDIR /go/src/morningo
EXPOSE 9000
#ENTRYPOINT /go/bin/morningo --port 9000
CMD ["/go/bin/morningo"]



