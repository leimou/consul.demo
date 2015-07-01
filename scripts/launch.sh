#!/bin/bash

# Create port file if not exist. The port file maintains the last port used
# to launch a service.
SEQ_FILE=$PWD/.lastport

if [ ! -f $SEQ_FILE ]; then
		echo "0" > $SEQ_FILE
fi

SEQ=$(head -1 $SEQ_FILE)

ARGC=$#
NUM=
if [ $ARGC -ge 1 ]; then
		NUM=$1
else
		NUM=1
fi

for i in $(seq 1 $NUM); do
		SEQ=$((SEQ+1))
		CONN_PORT=$((SEQ+10000))
		HTTP_PORT=$((SEQ+5000))
		echo "Launch service with port: $CONN_PORT/$HTTP_PORT"
		docker run -d --name fep.$SEQ -e SERIVCE_ID=fep.$SEQ -e SERVICE_32768_NAME=fep -e SERVICE_5000_NAME=fepinfo -p $CONN_PORT:32768 -p $HTTP_PORT:5000 consul.demo/fep
done

echo $SEQ > $SEQ_FILE
