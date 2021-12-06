#!/bin/sh

# Test GET, get one task
curl -X GET \
  --url http://localhost:8080/tasks/1 \
  -H 'cache-control: no-cache' \
  -w "\n get_one_http_code:%{http_code}\n"

# Test POST, create one task
curl -X POST \
  --url http://localhost:8080/tasks \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -w "\n post_http_code:%{http_code}\n" \
  -d '{"title":"create write code test post","dueDate":"06/12/2021","status":"todo"}'

# Test GET, get all task
curl -X GET \
  --url 'http://localhost:8080/tasks' \
  -H 'cache-control: no-cache' \
  -w "\n get_all_http_code:%{http_code}\n"

# Test GET, get expiring task
curl -X GET \
  --url 'http://localhost:8080/tasks?expiredToday=--expiring-today' \
  -H 'cache-control: no-cache' \
  -w "\n get_exp_http_code:%{http_code}\n"

# Test PUT, modify one task
curl -X PUT \
  --url http://localhost:8080/tasks \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -w "\n put_http_code:%{http_code}\n" \
  -d '{"id": 2,"title":"create write code test put","dueDate":"21/08/2019","status":"todo"}'

# Test DELETE, delete one task
curl -X DELETE \
  --url http://localhost:8080/tasks/12 \
  -H 'cache-control: no-cache' \
  -w "\n delete_http_code:%{http_code}\n"
