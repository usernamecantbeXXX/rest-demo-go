# rest-demo-go

# Restful API Demo

## 1. Build and Package From Source code

### 1.1 Run with Script

```
sh ./start.sh
sh ./stop.sh
sh ./restart.sh
```

### 1.1 Compile and Run

Compile and run the poject without script

```
sudo git clone https://github.com/usernamecantbeXXX/rest-demo-go.git
cd rest-demo-go/
# export GOPROXY=https://goproxy.cn,direct

go build
./rest-demo-go
```

## 2. unit test
```
./unit-test.sh
```
Test result is as below:
```shell
ubuntu@nmt:~/GolandProjects/rest-demo-go$ ./unit-test.sh 
{"id":2,"title":"create write code test put","dueDate":"21/08/2019","status":"todo"}
get_one_http_code:200

{"id":15,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"}
post_http_code:200

[{"id":1,"title":"write some code","dueDate":"18/11/2021","status":"done"},{"id":2,"title":"create write code test put","dueDate":"21/08/2019","status":"todo"},{"id":3,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":4,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":5,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":6,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":7,"title":"content","dueDate":"17/11/2021","status":"3"},{"id":8,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":9,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":10,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":11,"title":"content","dueDate":"17/11/2021","status":"done"},{"id":12,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":13,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":14,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":15,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"}]
get_all_http_code:200

[{"id":12,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":13,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":14,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"},{"id":15,"title":"create write code test post","dueDate":"06/12/2021","status":"todo"}]
get_exp_http_code:200

{"id":2,"title":"create write code test put","dueDate":"21/08/2019","status":"todo"}
put_http_code:200

delete_http_code:200

```

## 3. Call the API Command Line

After start the app, call the API with Command Line mode As Below：

```
# GET
sh ./tasks.sh list
sh ./tasks.sh list --expiring-today

# POST
sh ./tasks.sh add "curl add" "17/11/2021" "todo"

# PUT 
sh ./tasks.sh done 3 "curl put" "18/11/2021" "done"

# DELETE
sh ./tasks.sh delete 15

```

## 4. API Overview

| Method | HTTP Requests         | Returns          | Command                                 |
| ------ | --------------------- | ---------------- | --------------------------------------- |
| create | `POST /tasks`         | Created Task     | tasks add "write some code" 21/08/2019  |
| update | `PUT /tasks`          | Updated Task     | tasks done 3                            |
| delete | `DELETE /tasks/${id}` | Responce Code    | tasks delete 3                          |
| list   | `GET /tasks`          | A Task List      | tasks list /tasks list --expiring-today |

## 5. Architecture

### 5.1 Technical Stacks

- Go 1.16
- mux

### 5.2 Files Description

```
xxx@XXX:/mnt/d/xxx/GolandProjects/rest-demo-go$
.
├── README.md
├── generated-requests.http
├── go.mod
├── go.sum
├── log.out         # log file
├── main.go         # code
├── tasks.json      # json file as db for store the tasks
├── tasks.sh        # script for start the Restful web application
└── unit-test.sh    # unit Test

```