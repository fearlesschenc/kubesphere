#!/usr/bin/env bash

set -ex
set -o pipefail

# push to kubespheredev with default latest tag
TAG=${TAG:-latest}
# REPO=${REPO:-kubespheredev}
REPO=${REPO:-registry.cn-hangzhou.aliyuncs.com/fearlesschenc/containers/phoenix}

docker build -f build/ks-apiserver/Dockerfile -t $REPO/ks-apiserver:$TAG .
docker build -f build/ks-controller-manager/Dockerfile -t $REPO/ks-controller-manager:$TAG .

# Push image to dockerhub, need to support multiple push
cat ~/.docker/config.json | grep index.docker.io
if [[ $? != 0 ]]; then
  echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
fi
docker push $REPO/ks-apiserver:$TAG
docker push $REPO/ks-controller-manager:$TAG
