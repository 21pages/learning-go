package main

import (
	"geeorm"

	"geeorm/log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	e := geeorm.NewEngine("mysql", "sun:root@tcp(127.0.0.1:3306)/gee")
	defer e.Close()
	s := e.NewSession()
	s.Raw("drop table if exists User;").Exec()
	s.Raw("create table User (Name Text, `Age` integer);").Exec() //字段可以加``, 也可以不加
	s.Raw("create table User (Name Text, Age integer);").Exec()
	result, _ := s.Raw("insert into User (Name, Age) values ('zhangsan', 11), ('lisi', 12)").Exec()
	cnt, _ := result.RowsAffected()
	log.Info("rows:", cnt)
}

/*
compile:
	go build cmd_test/main.go

output:

[inf] 2022/01/06 15:05:29 geeorm.go:26: connect success! mysql sun:root@tcp(127.0.0.1:3306)/gee
[inf] 2022/01/06 15:05:29 raw.go:31: drop table if exists User;  []
[inf] 2022/01/06 15:05:29 raw.go:31: create table User (Name Text, Age integer);  []
[inf] 2022/01/06 15:05:29 raw.go:31: create table User (Name Text, Age integer);  []
[err] 2022/01/06 15:05:29 raw.go:33: Error 1050: Table 'User' already exists
[inf] 2022/01/06 15:05:29 raw.go:31: insert into User (`Name`, `Age`) values ('zhangsan', 11), ('lisi', 12)  []
[inf] 2022/01/06 15:05:29 main.go:20: rows: 2

sql:
mysql> select * from User;
+----------+------+
| Name     | Age  |
+----------+------+
| zhangsan |   11 |
| lisi     |   12 |
+----------+------+
2 rows in set (0.00 sec)

*/
