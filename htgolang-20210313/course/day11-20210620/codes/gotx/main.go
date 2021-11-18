/*
	事务
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	// 打开数据库连接池
	// dsn := "user:password@tcp(host:port)/dbName?charset=utf8mb4&parseTime=true"
	dsn := "root:9tUz2LK8wymRLz3@tcp(10.106.37.131:3306)/cmdb?charset=utf8mb4&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// 通常只有在程序退出时才关闭数据库连接
	defer db.Close()

	// 测试数据库的连接
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("truncate table alarm;")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sql := "insert into alarm(id,content,alarm_time) values(?,?,?);"
	//_, err = db.Exec(sql, 1, "第一条告警", time.Now())
	_, err = tx.Exec(sql, 1, "第一条告警", time.Now())
	fmt.Println(err)
	if err == nil {
		_, err = tx.Exec(sql, 2, "第二条告警", time.Now())
		fmt.Println(err)
	}

	if err == nil {
		// 所有 sql 都成功，才提交事务
		tx.Commit()
	} else {
		// 有一条 sql 不成功（有报错），就回退，sql 全部不生效
		tx.Rollback()
	}
}
