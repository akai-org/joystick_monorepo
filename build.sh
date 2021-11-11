#!/bin/bash
SERVICE_NAME=$1

docker-compose build $SERVICE_NAME
docker-compose up -d $SERVICE_NAME
if [ -z $SERVICE_NAME ]
then
	docker-compose logs -f
fi;
