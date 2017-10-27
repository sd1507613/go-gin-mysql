package main

import (
	"database/sql"
	"fmt"
    _"mysql"
	"github.com/gin-gonic/gin"
)

var (
	user        string = "root"
	password    string = "root"
	connectType string = "@tcp(localhost:3306)"
	database    string = "/zhengzhou?charset=utf8"
	connect     string = user + ":" + password + connectType + database
)

func main() {
	//insert()
	r := gin.Default()  
    r.GET("/ping", postDemoData)
	r.GET("/", sayHelloWorld)
	r.GET("/login", userlogin)    
    r.Run(":8001") // listen and server on 0.0.0.0:8001
}

func postDemoData(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{  
            "message": "pong",  
        })  
}

func sayHelloWorld(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.String(200, "Hello World")
}

//用户登录
func userlogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	
	username := c.Query("username")
	password := c.Query("password")
	
	fmt.Println(username)
	fmt.Println(password)
	
	db, err := sql.Open("mysql", connect)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		if  (record["username"] == username) &&
   		    (record["password"] == password){
				c.String(200, "true")
				return
			}
		fmt.Println(record)
	}
	c.String(200, "false")
}














/************************数据库demo**********************/

//插入demo
func insert() {
	db, err := sql.Open("mysql", connect)
	checkErr(err)

	stmt, err := db.Prepare(`INSERT tbl_app_report (case_id,XingMing,DiDian) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony", "tony", "tony")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//查询demo
func query() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	//普通demo
	//for rows.Next() {
	//	var userId int
	//	var userName string
	//	var userAge int
	//	var userSex int

	//	rows.Columns()
	//	err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//	checkErr(err)

	//	fmt.Println(userId)
	//	fmt.Println(userName)
	//	fmt.Println(userAge)
	//	fmt.Println(userSex)
	//}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
