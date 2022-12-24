package dataframe

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestFromSQL(t *testing.T) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/tsum_test")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	df, err := FromSQL(tx, "SELECT * FROM item_sku ORDER BY id LIMIT 10", []any{})

	fmt.Println(df, err)
}
