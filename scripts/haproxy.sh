#!/bin/bash

docker run -it --name consul-template --net=host -p 5000:5000 -p 8001:8001 --rm demo/haproxy
