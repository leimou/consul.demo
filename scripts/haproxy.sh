#!/bin/bash

docker run -it --name consul-template --net=host -p 5000:5000 -p --rm consul.demo/haproxy
