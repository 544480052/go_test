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
	db, error := sql.Open("mysql", "im:KGWkVR9NjmDxf0v2@tcp(10.105.62.216:3306)/im?charset=utf8")
	if error != nil {
		fmt.Println(error)
	}
	//Open 可能只是验证其参数而不创建与数据库的连接。要验证数据源名称是否有效，请调用 Ping
	error = db.Ping()
	if error != nil {
		fmt.Println(error)
	}

	//执行sql
	rows, error := db.Query("select staff_name from user")
	defer rows.Close() //方法执行结束之前关闭链接
	if error != nil {
		fmt.Println(error)
	}

	var name []string
	//name :=[...]string{}
	index :=0
	for rows.Next() {
		var staff_name string
		err := rows.Scan(&staff_name)
		if err != nil {
			fmt.Println(err)
		}
		name = append(name,staff_name)
		index++
	}
	fmt.Println(name)

	err := rows.Err()
	if err != nil {
		fmt.Println(err)
	}

}
