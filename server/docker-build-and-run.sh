#! /bin/bash

currentDir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

docker rm -f mymongo
docker run --name mymongo -p 27017:27017  -d mongo:3.4

IMAGE_NAME="comments-server:latest"

docker build -t $IMAGE_NAME  $currentDir

docker run -p 3000:3000 --link mymongo:mymongo.org $IMAGE_NAME
