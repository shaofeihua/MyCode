package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 打开数据库连接池
	// dsn := "user:password@tcp(host:port)/dbName?charset=utf8mb4&parseTime=true"
	dsn := "root:9tUz2LK8wymRLz3@tcp(10.106.37.131:3306)/cmdb?charset=utf8mb4&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//sql := "select id,content,alarm_time from alarm limit 0"
	sql := "select * from user limit 0"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	//关闭 rows ，告知将连接放入连接池（连接使用完了，将连接释放）
	defer rows.Close()

	columnTypes, _ := rows.ColumnTypes()
	columns, _ := rows.Columns()

	for i, typ := range columnTypes {
		fmt.Println(i, columns[i])
		fmt.Println(typ.DatabaseTypeName())
		fmt.Println(typ.Length())
		fmt.Println(typ.Nullable())
		fmt.Println(typ.ScanType())
		fmt.Println(typ.DecimalSize())
		fmt.Println()
	}
}
