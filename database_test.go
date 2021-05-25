package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root-pass@tcp(localhost:33060)/db_golang_belajar_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
