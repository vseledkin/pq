package pq

import (
	"testing"
	"database/sql"
)

func TestKerberos(t *testing.T) {

	connStr := "user=unstructapisvc@PRIVATE dbname=nsidb_prognoz sslmode=disable host=is27009.private krbsrvname=is27009.private"
	db, e := sql.Open("postgres", connStr)
	if e != nil {
		t.Fatal(e)
	}

	row := db.QueryRow("SELECT 1")
	var num int
	e = row.Scan(&num)
	if e != nil {
		t.Fatal(e)
	}

	if num != 1 {
		t.Fatalf("Value failed, want 1 got %+v", num)
	}

}
