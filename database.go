package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbCredentials struct {
	client   string
	user     string
	password string
	port     string
	host     string
	database string
}

func (db *dbCredentials) formatStr() string {
	return db.client + "://" + db.user + ":" + db.password + "@" + db.host + ":" + db.port + "/" + db.database
}

func Connectdb() string {
	fmt.Println("hello")
	sql.Drivers()
	// creds := dbCredentials{
	// 	client: "postgresql", user: "postgres",
	// 	password: "postgresql", port: "5432",
	// 	host:     "localhost",
	// 	database: "teste",
	// }
	db, err := sql.Open("postgres", "user=postgres host=localhost port=5432 password=postgresql dbname=teste sslmode=disable")
	if err != nil {
		println(err)
	}
	rows, err := db.Query("select * from lul ")
	if err != nil {
		panic(err)
	}
	var id int
	var name string
	rows.Scan(&id, &name)
	println("connected to db")
	return "connected to db"
}
