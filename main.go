package main

import (
	"database/sql"
	//go mysql驱动需要下载【下载：go get github.com/Go-SQL-Driver/MySQL 安装：go install github.com/Go-SQL-Driver/MySQL】
	//github.com会安装在$GOPATH的src目录下
	_ "github.com/Go-SQL-Driver/MySQL"
	"fmt"
)

func main() {
	//链接数据库【驱动，数据库信息（用户名:密码@tcp(IP:端口)/数据库名?charset=utf8）】
	db, error := sql.Open("mysql", "root:''@tcp(127.0.0.1:3306)/dev_im?charset=utf8")
	if error != nil {
		fmt.Println(error)
	}
	//Open 可能只是验证其参数而不创建与数据库的连接。要验证数据源名称是否有效，请调用 Ping
	error = db.Ping()
	if error != nil {
		fmt.Println(error)
	}

	//执行sql
	rows,error:= db.Query("select * from users")
	defer rows.Close()//方法执行结束之前关闭链接
	if error!=nil {
		fmt.Println(error)
	}

	for rows.Next(){
		var staff_name string
		err:= rows.Scan(&staff_name)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(staff_name)
	}

	err := rows.Err()
	if err!=nil {
		fmt.Println(err)
	}


}
