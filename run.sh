#!/usr/bin/env bash
tag=$(git log -1 --pretty=%h)

go build
docker build -t $tag --build-arg SHA=$tag .
docker run -p 8081:8081 $tag
 
# the sha tag is useful for ehen its deployed , ci/cd 
