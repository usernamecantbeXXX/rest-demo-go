#!/bin/sh

host="127.0.0.1"
port="8080"

if [ "$1" = "list" -a "$2" = "--expiring-today" ];then
  curl "http://$host:$port/tasks?expiredToday=--expiring-today"
  echo "expiring today"
  exit 0
elif [ "$1" = "list" -a -z "$2" ];then
  curl "http://$host:$port/tasks"
  echo "list"
  exit 0
elif [ "$1" = "add" ];then

  if [ -z "$2" -a -z "$3" -a -z "$4" ]; then
      echo "Task param all null, failed"
      exit 2
  fi

  curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'title='"${2}"'&dueDate='"${3}"'&status='"${4}"'' "http://$host:$port/tasks"
  echo "add "
  exit 0

elif [ "$1" = "done" ];then

  if [ -z "$2" -a -z "$3" -a -z "$4" ]; then
      echo "Task param all null, failed"
      exit 2
  fi

  curl -X PUT -H "Content-Type: application/json" -d '{"id":"'"${2}"'","title":"'"${3}"'","dueDate":"'"${4}"'","status":"'"${5}"'"}' "http://$host:$port/tasks"
  echo "update "
  exit 0

elif [ "$1" = "delete" ];then

    curl -X DELETE "http://$host:$port/tasks/${2}"
    echo "delete "
    exit 0

fi