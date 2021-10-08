#!/usr/bin/env bash
tag=$(git log -1 --pretty=%h)
time=$(date)

go build 
docker build -t "$tag" --build-arg SHA="$tag" --build-arg TIME="$time" .
docker run -p 8081:8081 "$tag"
 

#script to terminate when the build fails