#!/bin/bash

# Build docker image
#docker-compose up -d --build
# Run all the services except the consumers, which will be run after all the apps are prepaired
docker-compose up -d --build

# kill -INT $(cat pid) && ./morningo # graceful stop the process and restart


# go get github.com/chenhg5/morningo-installer
 
docker start go_db 
docker start go_web
docker start go_nginx