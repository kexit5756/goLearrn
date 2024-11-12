package main

import "fmt"

func insertRowDemo() {
	sqlstr := "insert into user(name,age) values(?,?)"
	ret, err := db.Exec(sqlstr, "王五", 18)
	if err != nil {
		fmt.Println("Insert failed")
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("get LastInsertId failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
func updateDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
