#!/bin/bash

echo "Stopping running fep containers ..."
docker stop $(docker ps -a | grep -o "fep\.[0-9][0-9]*")

echo "Cleaning up fep containers ..."
docker rm $(docker ps -a | grep -o "fep\.[0-9][0-9]*")
