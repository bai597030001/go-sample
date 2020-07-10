package main

// 1.下载并导入数据库驱动包
// go get github.com/go-sql-driver/mysql
import (
	"container/list"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化并返回一个*sql.DB结构体对象,sql.DB结构是sql/database包封装的一个数据库操作对象（数据库连接池），包含了操作数据库的基本方法。
// func Open(driverName, dataSourceName string) (*DB, error) // 一般而言，我们使用Open()函数便可初始化并返回一个*sql.DB结构体实例
// func OpenDB(c driver.Connector) *DB // OpenDB()函数则依赖驱动包实现sql/database/driver包中的Connector接口，这种方法并不通用化

// 定义一个全局对象db
var db *sql.DB
var err error

type User struct {
	id         int64
	username   string
	passwd     string
	email      string
	createtime string
	updatetime string
}

func getConnect() *sql.DB {
	// 2.连接数据库
	db, err = sql.Open("mysql", "root:gsmrlab@tcp(localhost:3306)/test")
	if db == nil || err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// 设置数据库连接池属性
	db.SetMaxOpenConns(100)                //设置最多打开100个数据连连接
	db.SetMaxIdleConns(1)                  //设置为0表示不限制打开连接数
	db.SetConnMaxLifetime(time.Second * 5) //5秒超时
	return db
}

func getSingleData() *User {
	selectText := "SELECT * FROM user where id = ?;"
	var user User
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(selectText, 2).Scan(&user.id, &user.username, &user.passwd, &user.email, &user.createtime, &user.updatetime)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	fmt.Println(user)
	return &user
}

func getMiltiData() *list.List {
	// 3.执行查询
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	// func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
	// Query()和QueryContext()方法返回一个sql.Rows结构体，代表一个查询结果集
	selectText := "SELECT * FROM user;"
	rows, err := db.Query(selectText)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 读取数据
	l := list.New()
	for rows.Next() {
		var (
			id         int64
			username   string
			passwd     string
			email      string
			createtime string
			updatetime string
		)
		_ = rows.Scan(&id, &username, &passwd, &email, &createtime, &updatetime)
		fmt.Println(id, username, passwd, email, createtime, updatetime)
		l.PushBack(username)
	}
	return l
}

func insertData() int64 {
	sqlStr := "insert into user(username, passwd, email) values (?,?,?)"
	ret, err := db.Exec(sqlStr, "adm", "123456", "123@gmail.com")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0
	}
	// 新插入数据的id
	affectedRow, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0
	}
	return affectedRow
}

func updateData() int64 {
	sqlStr := "update user set passwd = ? where username = ?"
	ret, err := db.Exec(sqlStr, "asd", "adm")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return 0
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return 0
	}
	fmt.Printf("update success, affected rows:%d\n", n)
	return n
}

func deleteData() int64 {
	sqlStr := "delete from user where username = ?"
	ret, err := db.Exec(sqlStr, "adm")
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return 0
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return 0
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
	return n
}

func dbOperate() {
	getConnect()
	getSingleData()
	getMiltiData()
	insertData()
	updateData()
	deleteData()

	prepareQuery()
	prepareInsert()
	prepareDelete()
}

/**
# MySQL预处理

普通SQL语句执行过程：
	客户端对SQL语句进行占位符替换得到完整的SQL语句。
	客户端发送完整SQL语句到MySQL服务端
	MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：
	把SQL语句分成两部分，命令部分与数据部分。
	先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
	然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
	MySQL服务端执行完整的SQL语句并将结果返回给客户端。

## 为什么要预处理？
	优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
	避免SQL注入问题
*/

func prepareQuery() {
	sqlStr := "select * from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.username, &user.passwd, &user.email, &user.createtime, &user.updatetime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d username:%s passwd:%s email:%s createtime: %s updatetime:%s\n", user.id, user.username, user.passwd, user.email, user.createtime, user.updatetime)
	}
}

func prepareInsert() {
	sqlStr := "insert into user(username, passwd, email) values (?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("adm", "123456", "123@gmail.com")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("adm1", "123456", "123@gmail.com")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

func prepareDelete() {
	sqlStr := "delete from user where username = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("adm")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("adm1")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("delete success.")
}

/**
## 事务相关方法
Go语言中使用以下三个方法实现MySQL中的事务操作。

开始事务
func (db *DB) Begin() (*Tx, error)

提交事务
func (tx *Tx) Commit() error

回滚事务
func (tx *Tx) Rollback() error
*/
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set passwd='abc' where id=?"
	_, err = tx.Exec(sqlStr1, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update user set passwd='def' where id=?"
	_, err = tx.Exec(sqlStr2, 4)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}
