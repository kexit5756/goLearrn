package main

import "fmt"

// Prepare预处理操作
func prepareQueryDemo() {
	sqlstr := "select id,name,age from user where id>?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("Prepare failed,err%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("quer falied err%v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan falied %v\n", err)
		}
		fmt.Printf("id%d,name%s,age%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func prepareInsertDemo() {
	sqlstr := "insert into user(name, age) values (?,?)"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed%v\n")
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("刘3", 18)
	if err != nil {
		fmt.Printf("Insert falied")
		return
	}
	_, err = stmt.Exec("刘2", 30)
	if err != nil {
		fmt.Printf("Insert falied")
		return
	}
	fmt.Printf("InsertSucess\n")
}
