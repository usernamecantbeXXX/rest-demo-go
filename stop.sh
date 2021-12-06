#!/bin/sh
PIDFILE=service.pid
if [ -f "$PIDFILE" ]; then
     kill -15 `cat $PIDFILE`
     rm -rf $PIDFILE
     echo "Service is stop SUCCESS!"
else
    echo "Service is already stop ..."
fi