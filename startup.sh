#!/bin/bash

# Build docker image
#docker-compose up -d --build
# Run all the services except the consumers, which will be run after all the apps are prepaired
docker-compose up -d --build

# kill -INT $(cat pid) && ./app # graceful stop the process and restart
 
docker start go_v1_db 
docker start go_v1_web
docker start go_v1_nginx