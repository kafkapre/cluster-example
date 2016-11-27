#! /bin/bash

currentDir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

IMAGE_NAME="angular:latest"

docker build -t $IMAGE_NAME  $currentDir

docker run -p 3000:3000 $IMAGE_NAME
