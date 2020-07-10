package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var sqlxDb *sqlx.DB

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/test"
	// 也可以使用MustConnect连接不成功就panic
	sqlxDb, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	sqlxDb.SetMaxOpenConns(20)
	sqlxDb.SetMaxIdleConns(10)
	return
}

func queryRow() {
	sqlStr := "select id, name, age from user where id=?"
	var user User
	err := sqlxDb.Get(&user, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s passwd:%s\n", user.id, user.username, user.passwd)
}

func queryMultiRow() {
	sqlStr := "select * from user where id > ?"
	var users []User
	err := sqlxDb.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

func insertRow() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := sqlxDb.Exec(sqlStr, "沙河小王子", 19)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRow() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := sqlxDb.Exec(sqlStr, 39, 6)
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
func deleteRow() {
	sqlStr := "delete from user where id = ?"
	ret, err := sqlxDb.Exec(sqlStr, 6)
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

// ## 事务操作
// sqlx中提供的db.Beginx() 和 tx.MustExec() 方法来简化错误处理过程。
func transaction() {
	tx, err := sqlxDb.Beginx() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=40 where id=?"
	tx.MustExec(sqlStr1, 2)
	sqlStr2 := "Update user set age=50 where id=?"
	tx.MustExec(sqlStr2, 4)
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}