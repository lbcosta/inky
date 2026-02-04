package db

import (
	"database/sql"
	"fmt"
	"math/rand/v2"

	_ "github.com/tursodatabase/turso-go"
)

func connect() (*sql.DB, error) {
	conn, err := sql.Open("turso", ":memory:")
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func Get() (string, error) {
	conn, err := connect()
	if err != nil {
		fmt.Printf("Error on db.Get: %v\n", err)
		return "", err
	}
	defer conn.Close()

	sql := "CREATE table go_turso (foo INTEGER, bar TEXT)"
	_, _ = conn.Exec(sql)

	sql = "INSERT INTO go_turso (foo, bar) values (?, ?)"
	stmt, _ := conn.Prepare(sql)
	defer stmt.Close()

	n1 := rand.N(999)
	// generate a random uppercase letter A-Z
	n2 := string(rune(rand.N(26) + 65))

	_, _ = stmt.Exec(n1, n2)
	rows, _ := conn.Query("SELECT * from go_turso")
	defer rows.Close()

	result := ""
	for rows.Next() {
		var a int
		var b string
		_ = rows.Scan(&a, &b)
		result = fmt.Sprintf("%s%d", b, a)
	}

	return result, nil
}
