#!/bin/bash

ARGC=$#


while (( $# )); do
		if [ $1 = "restart" ]; then
				echo "==> Note: Cluster meta data deleted"
				rm -rf ~/consul
		elif [ $1 == "server" ]; then
				echo "==> Note: Launching consul server in bootstrap mode"
				consul agent -server -data-dir="/home/manage/consul" -bootstrap-expect 3 $@
				break
		else
				echo "==> Note: Launching consul client"
				consul agent -data-dir="/home/manage/consul" -join 192.168.16.204 192.168.16.205 192.168.16.206 $@
	  fi
		shift
done
