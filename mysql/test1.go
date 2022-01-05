package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Name string `db:"Name"`
	Age  int    `db:"Age"`
}

func connect(driver, source string) *sqlx.DB {
	db, err := sqlx.Open(driver, source)
	if err != nil {
		log.Println("open:", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Println("ping:", err) // not really connect
		db.Close()
		return nil
	}
	return db
}

func insert(db *sqlx.DB, users []User) error {
	if len(users) == 0 {
		return fmt.Errorf("insert len = 0")
	}
	query := "insert into User (Name, Age) values "
	for i := 0; i < len(users); i++ {
		query += fmt.Sprintf("(\"%s\", %d)", users[i].Name, users[i].Age) //字符串注意加\"
		if i != len(users)-1 {
			query += ","
		}
	}
	//query += ";" // not necessary
	_, err := db.Exec(query)
	if err != nil {
		log.Println("insert:", err)
		return err
	}

	return nil
}

/*
   1) 原子性
   2) 一致性
   3) 隔离性
   4) 持久性
*/
func insertTransaction(db *sqlx.DB, users []User) error {
	if len(users) == 0 {
		return fmt.Errorf("insert user num = 0")
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("insert begin:", err)
		return err
	}
	for i := 0; i < len(users); i++ {
		query := fmt.Sprintf("insert into User (Name, Age) values (\"%s\", %d)", users[i].Name, users[i].Age)
		if _, err := tx.Exec(query); err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	}
	return tx.Commit()
}

func search(db *sqlx.DB, query string) ([]User, error) {
	user := []User{}
	if err := db.Select(&user, query); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func delete(db *sqlx.DB, user User) error {
	if _, err := db.Exec(fmt.Sprintf("delete from User where Name=\"%s\" and Age=%d", user.Name, user.Age)); err != nil {
		log.Println("delete:", err)
		return err
	}
	return nil
}

func update(db *sqlx.DB, old User, new User) error {
	if _, err := db.Exec(fmt.Sprintf("update User set Name=\"%s\", Age=%d where Name=\"%s\" and Age=%d", new.Name, new.Age, old.Name, old.Age)); err != nil {
		log.Print("update:", err)
		return err
	}
	return nil
}

func main() {
	db := connect("mysql", "sun:root@tcp(127.0.0.1:3306)/gee")
	if db == nil {
		return
	}
	defer db.Close()
	insert(db, []User{{"Tom", 11}, {"Sam", 12}})
	insertTransaction(db, []User{{"Peter", 13}, {"Jone", 14}})
	update(db, User{"Tom", 11}, User{"Tommy", 15})
	delete(db, User{"Peter", 13})
	users, err := search(db, "select * from User where Age > 10 and Age < 20")
	if err == nil {
		log.Println(users)
	}
}

/*
output:
2022/01/05 16:43:35 [{Tommy 15} {Sam 12} {Jone 14}]
*/
