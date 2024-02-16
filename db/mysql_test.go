package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("create table if not exists `test`(`id` int)")
	if err != nil {
		fmt.Println("create table failed", err)
		return
	}

	_, err = db.Exec("insert into `test` values (2024)")
	if err != nil {
		fmt.Println("insert failed", err)
		return
	}

	result, err := db.Query("select * from `test`")
	if err != nil {
		fmt.Println("query failed", err)
		return
	}
	for result.Next() {
		id := 0
		err = result.Scan(&id)
		if err != nil {
			fmt.Println("scan failed", err)
			return
		}
		fmt.Println("id is", id)
	}

	_, err = db.Exec("delete from `test`")
	if err != nil {
		fmt.Println("delete failed", err)
		return
	}

	_, err = db.Exec("drop table `test`")
	if err != nil {
		fmt.Println("drop table failed", err)
		return
	}
}
