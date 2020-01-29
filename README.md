# Go-Lemonilo
Golang CRUD Lemonilo

## Installation
```
1. go get "github.com/astaxie/beego"
2. go get "golang.org/x/crypto/bcrypt"
3. go get "github.com/go-sql-driver/mysql"
4. go get "github.com/astaxie/beego/orm"
```

## Database
```
" create or replace table users
(
  id int (6) unsigned auto_increment
  primary key,
  name varchar (30) not null,
  email varchar (30) not null,
  address text null,
  password varchar (250) not null
); "
```
## Config

update config at conf/app.conf with your database


## Running
```
go run main.go
```


