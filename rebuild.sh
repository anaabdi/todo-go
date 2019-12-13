#!/bin/bash

if ! [ -x "$(command -v docker-compose)" ]; then
  echo "Error: docker-compose is not installed"
  exit 1
fi

docker-compose stop $@
docker-compose rm -f $@
docker-compose build $@
docker-compose up -d $@
docker images | grep '<none>' | awk '{ print $3 }' | xargs docker rmi || true