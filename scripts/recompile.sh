#!/bin/bash

docker exec -it go_web go install github.com/kardianos/govendor

docker exec -it go_web make deps
docker exec -it go_web go build -o /go/bin/morningo -tags=jsoniter -v ./
docker restart go_web