# GoAPI_Test
SQL 設定
``` sql
CREATE DATABASE test;

USE test;

CREATE TABLE user_account(
    id VARCHAR(10) NOT NULL,
    uname VARCHAR(10) character set UTF8  NOT NULL, 
    memo varchar(100) character set UTF8  NOT NULL, 
    primary key (id));


insert into user_account VALUE('xxxxx', 'idname', 'test');  
```

Go API Server 設定
``` bash 
$ go run .
$ curl http://localhost:8080/testGet
```
