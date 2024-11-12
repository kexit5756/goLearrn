package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb() (err error) {
	user := "root"
	password := "123456"
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/"
	db, err := sql.Open("mysql", dsn)
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// func (db *DB) SetMaxOpenConns(n int)
// PS:设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，
// 会将最大闲置连接数减小到匹配最大开启连接数的限制。
// 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
func main() {
	// DSN:Data Source Name
	err := InitDb()
	if err != nil {
		fmt.Println("initDb err:", err)
		return
	}
	fmt.Println("initDb success")
}
