# day1 sql基础

## 创建mysql数据库
1. sudo apt-get install mysql-server mysql-client
2. mysql报错`mysql ERROR 1045 (28000): Access denied for user 'sun'@'localhost' (using pa`, 初始用户及密码在`/etc/mysql/debian.cnf`, 这里使用在`/etc/mysql/mysql.conf.d/mysqld.cnf`的`[mysqld]`分组中添加`skip-grant-tables`, 然后`service mysql restart`.
3. `create database gee;`
4. `use gee;`
5. `create table User(Name text, Age integer);`
6. `insert into User (Name, Age) value ("zhangsan", 27), ("lisi", 28)`

## 构建
1. session中使用database/sql, 来做通用的sql数据库访问, 使用的部分import mysql,sqlite即可
2. 自定义log输出
3. 定义engine
