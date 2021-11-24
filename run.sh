#!/usr/bin/env bash
tag=$(git log -1 --pretty=%h)
time=$(date)

go build 
docker build -t "$tag" --build-arg SHA="$tag" --build-arg TIME="$time" . || { echo "failed to build docker image"; exit 1; }
docker run -p 8081:8081 "$tag" -d \
  -it \
  --name devtest \
  --mount type=bind,source="$(pwd)"/target,target=/app \
  nginx:latest
 

#script to terminate when the build fails