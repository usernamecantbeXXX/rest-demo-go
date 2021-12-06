#!/bin/sh
PIDFILE=service.pid
if [ ! -d "logs" ]; then
   sudo mkdir logs
fi
if [ -f "$PIDFILE" ]; then
    echo "Service is already start ..."
else
    echo "Service  start ..."
    nohup go build && ./rest-demo-go 1> logs/rest_demo.out 2>&1  &
    printf '%d' $! > $PIDFILE
    echo "Service  start SUCCESS "
fi