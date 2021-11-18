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
	dsn := "golang:9tUz2LK8wymRLz3@tcp(10.106.37.131:3306)/cmdb?charset=utf8mb4&parseTime=true"
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

	fmt.Println("ok")

	fmt.Println("ddl: ", ddl(db))

	// dql
	rs, err := db.Exec("select * from alarm")
	fmt.Println(err)
	fmt.Println(rs) // {0xc000188000 0xc000182050}

	ddl(db)
	dmlInsert(db)
	dmlUpdate(db)
	dqlRows(db)
	dqlRow(db)
}

func ddl(db *sql.DB) error {
	// 执行
	sql := `
	create table if not exists alarm(
	  id bigint primary key auto_increment,
	  alarm_time datetime not null,
	  content text
	)engine=innodb default charset=utf8mb4;
	`
	_, err := db.Exec(sql)
	return err
}

func dmlInsert(db *sql.DB) error {
	// dml
	//拼接 sql
	alarmTime, alarmContent := time.Now().Format("2006-01-02 15:04:05"), "MEM告警"
	sql := "insert into alarm(alarm_time,content) values('" + alarmTime + "','" + alarmContent + "')"
	//rs, err := db.Exec("insert into alarm(alarm_time,content) values('2021-11-14 10:58:11','CPU告警')")
	rs, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	// 参数化 sql 。防止 sql 注入
	alarmTime, alarmContent = time.Now().Format("2006-01-02 15:04:05"), "防止 sql 注入"
	sql = "insert into alarm(alarm_time,content) values(?,?);"
	rs, err = db.Exec(sql, alarmTime, alarmContent)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())

	sql = "update alarm set content=?,status=? where id=?;"
	rs, err = db.Exec(sql, "修改 content", 100, 1)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())
	return nil
}

func dmlUpdate(db *sql.DB) error {
	sql := "update alarm set status=status+1 where id%2=0;"
	rs, err := db.Exec(sql)
	if err != nil {
		return err
	}
	fmt.Println(rs.LastInsertId())
	fmt.Println(rs.RowsAffected())
	return nil
}

func dqlRows(db *sql.DB) error {
	//sql := "select id,content,alarm_time from alarm where id=?;"
	//sql := "select id,content,alarm_time from alarm where content like ?"
	//sql := "select id,content,alarm_time from alarm where content like ? limit ?"
	sql := "select id,content,alarm_time from alarm where content like ? limit ? offset ?"

	rows, err := db.Query(sql, "%sql 注入%", 2, 5)
	if err != nil {
		return err
	}

	//关闭 rows ，告知将连接放入连接池（连接使用完了，将连接释放）
	defer rows.Close()

	for rows.Next() {
		var id int64
		var content string
		var alarmTime time.Time
		// 将扫描到的数据，赋值到变量中，变量的数量、顺序和数据类型，需要跟扫描结果中的数量、顺序、类型保持一致
		if err := rows.Scan(&id, &content, &alarmTime); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(id, content, alarmTime)
		}
	}

	return nil
}

func dqlRow(db *sql.DB) error {
	sql := "select id,content,alarm_time from alarm where id = ?"

	var id int64
	var content string
	var alarmTime time.Time

	if err := db.QueryRow(sql, 18).Scan(&id, &content, &alarmTime); err != nil {
		//fmt.Println(err)
		return err
	}
	fmt.Println(id, content, alarmTime)
	return nil
}
