#! /bin/bash

id=$(docker ps -a -q -f ancestor=angular:latest) 


docker rm -f $id