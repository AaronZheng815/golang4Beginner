package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //使用sqlx需要这行
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
type Place struct {
	Country  string `db:"country"`
	City     string `db:"city"`
	PostCode int    `db:"post_code"`
}

// var schema = `
// CREATE TABLE person (
//     user_id text,
//     user_name text,
// 	sex ,
// 	email text
// );

// CREATE TABLE place (
//     country text,
//     city text NULL,
//     telcode integer
// )`

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@(127.0.0.1:3306)/testdb")
	if err != nil {
		//fmt.Println("Open mysql failed. Error: ", err)
		log.Fatalln(err)
		return
	}
	Db = database
}

func main() {
	//insert a record
	r, err := Db.Exec("insert into person(user_name, sex, email) values(?,?,?)",
		"aaron zheng", "male", "aaron@qq.com")
	if err != nil {
		log.Fatalln("insert error:", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("insert Success: ", id)

	//select a record
	var person []Person
	err = Db.Select(&person, "select user_id,user_name,sex,email from person where user_name=?", "aaron zheng")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("select success: ", person)

	//update record
	_, err = Db.Exec("update person set user_name=? where user_id=?", "nina", 2)
	if err != nil {
		log.Fatalln("update failed: ", err)
		return
	}

	//delete record
	_, err = Db.Exec("delete from person where user_id=?", 1)
	if err != nil {
		log.Fatalln("delete error:", err)
		return
	}
	fmt.Println("delete sucess")
}
