// +build go1.10

package pq

import (
	"database/sql"
	"fmt"
)

func ExampleNewConnector() {
	name := ""
	connector, err := NewConnector(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	db := sql.OpenDB(connector)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Use the DB
	txn, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	txn.Rollback()
}
