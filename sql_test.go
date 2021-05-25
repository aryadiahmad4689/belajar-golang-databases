package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExcecSql(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id,name) values ('adi','Adi')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert to customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "Select * from customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	// fmt.Println(rows)
	defer rows.Close()

}

func TestQueryRows(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "Select * from customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id", id)
		fmt.Println("Name", name)

	}
	// fmt.Println(rows)
	defer rows.Close()
}

func TestQueryRowsComplex(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "Select id,name,email,balance,bhirtday,rating,created_at,married from customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var bhirtday, createdAt time.Time
		var married bool

		var err = rows.Scan(&id, &name, &email, &balance, &bhirtday, &rating, &createdAt, &married)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id", id)
		fmt.Println("Name", name)
		if email.Valid {
			fmt.Println("email", email.String)
		}
		fmt.Println("balance", balance)
		fmt.Println("rating", rating)
		fmt.Println("bhirtday", bhirtday)
		fmt.Println("created at", createdAt)
		fmt.Println("married", married)

	}
	// fmt.Println(rows)
	defer rows.Close()
}

func TestSqlAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	comment := "halo ini sangat bagus"
	result, err := db.ExecContext(ctx, "Insert into comments(comment) values(?)", comment)

	if err != nil {
		panic(err)
	}

	lastInsert, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsert)

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	state, err := db.PrepareContext(ctx, "Insert into comments(comment) values(?)")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		comment := "halo ini sangat bagus ke " + strconv.Itoa(i)
		result, err := db.ExecContext(ctx, "Insert into comments(comment) values(?)", comment)

		if err != nil {
			panic(err)
		}

		lastInsert, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println(lastInsert)
	}
	defer state.Close()

}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	state, err := tx.PrepareContext(ctx, "Insert into comments(comment) values(?)")

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		comment := "halo ini sangat bagus ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, "Insert into comments(comment) values(?)", comment)

		if err != nil {
			panic(err)
		}

		lastInsert, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println(lastInsert)
	}

	defer state.Close()

	tx.Rollback()
}
