package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

func createLine(str string) {
	fmt.Printf("=================== %s ===================\n", str)
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "go_basic_user:go_basic_password@tcp(127.0.0.1:3306)/go_basic_db"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	PersonID  int
	FirstName string
	LastName  string
	Address   string
	City      string
}

// CREATE TABLE go_user (
// 	PersonID INT,
// 	LastName VARCHAR ( 255 ),
// 	FirstName VARCHAR ( 255 ),
// 	Address VARCHAR ( 255 ),
// City VARCHAR ( 255 )
// );

// insert data
func insertData() {
	sqlStr := "insert into go_user(PersonID, FirstName,LastName,Address,City) values (?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, 1, "zhang", "san", "beijing", "beijing")
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

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select PersonID, FirstName, LastName, Address, City from go_user where PersonID=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.PersonID, &u.FirstName, &u.LastName, &u.Address, &u.City)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d FirstName:%s LastName:%s Address:%s City:%s\n", u.PersonID, u.FirstName, u.LastName, u.Address, u.City)
}
func main() {
	createLine("START")
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	queryRowDemo()
}
