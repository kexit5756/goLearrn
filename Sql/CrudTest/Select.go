package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	age  int
	name string
}

var db *sql.DB

func InitDB() (err error) {
	user1 := "root"
	password := "123456"
	dsn := user1 + ":" + password + "@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	defer db.Close()
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id >= ?"
	var u user
	rows, err := db.Query(sqlStr, 1)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	//insertRowDemo()
	//updateDemo()
	prepareInsertDemo()
	prepareQueryDemo()

}
