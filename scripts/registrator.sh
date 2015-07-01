#!/bin/bash

docker run --net=host -v /var/run/docker.sock:/tmp/docker.sock gliderlabs/registrator consul:
