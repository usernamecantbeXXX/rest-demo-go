#!/bin/sh

curl --request GET \
  --url http://localhost:8080/tasks/1 \
  --header 'cache-control: no-cache' \

curl --request GET \
  --url http://localhost:8080/tasks \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --data '{\n	"title": "create write code",\n	"dueDate": "21/08/2019",\n	"status": "todo"\n}'

curl --request GET \
  --url 'http://localhost:8080/tasks?expiredToday=--expiring-today' \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --header 'postman-token: 1bf97590-0581-0982-e870-5b50aa3c24fc' \
  --data '{\n	"title": "create write code",\n	"dueDate": "21/08/2019",\n	"status": "todo"\n}'

curl --request PUT \
  --url http://localhost:8080/tasks \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --header 'postman-token: a8b2a4ba-1c70-9309-83a3-61f39ec99b09' \
  --data '{\n	"id": "2",\n	"title": "put write code",\n	"dueDate": "21/08/2019",\n	"status": "done"\n}'

curl --request DELETE \
  --url http://localhost:8080/tasks/1 \
  --header 'cache-control: no-cache' \
