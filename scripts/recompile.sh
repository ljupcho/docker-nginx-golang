#!/bin/bash

docker exec -it go_v1_web go install github.com/kardianos/govendor

# docker exec -it go_v1_web go mod download
docker exec -it go_v1_web go build -o /go/bin/app -tags=jsoniter -v ./
docker restart go_v1_web