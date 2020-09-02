#!/bin/bash

# Build docker image
#docker-compose up -d --build
# Run all the services except the consumers, which will be run after all the apps are prepaired
docker-compose up -d --build

# as this is map to the applicatio folder that has Makefile ex. with morningo
docker exec -it go_web make deps
# docker exec -it go_web make test
# docker exec -it go_web make restart
# docker exec -it go_web make
docker exec -it go_web make build
docker exec -it go_web make run
# docker exec -it go_web "CMD cd /go/src/morningo"

# kill -INT $(cat pid) && ./morningo # graceful stop the process and restart


# go get github.com/chenhg5/morningo-installer