package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 查询sql
func main() {

	db, err := sql.Open("mysql", "mi:4ZKk6SaCtLlGyIZWOBdRr1yAF1HoxfLGzZ@tcp(tj1-owt-mitob-staging-db-01.kscn:3306)/mitob_platform?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from sys_account")

	for rows.Next() {
		fmt.Println("dd")
		//row.Scan(...)
	}
	rows.Close()
}
