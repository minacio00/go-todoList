package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

type dbCredentials struct {
	client   string
	user     string
	password string
	port     string
	host     string
	database string
	ssl      string
}

func (db *dbCredentials) formatStr() string {
	return "user=" + db.user + " host=" + db.host + " port=" + db.port + " password=" + db.password + " dbname=" + db.database + " sslmode=" + db.ssl
} // "user=postgres host=192.168.1.2 port=5432 password=postgresql dbname=teste sslmode=disable"

func GetAllluls() map[int]string {
	db := connectdb()
	rows, err := db.Query("select * from lul ")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string
	var m = make(map[int]string)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		m[id] = name
		println(name, id)
	}
	return m

}

func GetTodos() map[string][]string {
	db := connectdb()
	rows, err := db.Query("select * from todos")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var id int
	var name string
	var description string
	var m = make(map[string][]string)

	println("results of todos query: ")
	for rows.Next() {
		err = rows.Scan(&id, &name, &description)
		if err != nil {
			panic(err)
		}
		m[name] = []string{strconv.Itoa(id), name, description}
		println(name, id, description)
	}
	return m

}

func connectdb() *sql.DB {
	fmt.Println("hello")
	sql.Drivers()
	creds := dbCredentials{
		client: "postgresql", user: "postgres",
		password: "postgresql", port: "5432",
		host:     "db", // "192.168.1.2",
		database: "teste",
		ssl:      "disable",
	}
	ch := make(chan *sql.DB)
	go func() {
		db, err := sql.Open("postgres", creds.formatStr())
		if err != nil {
			println(err)
		}
		ch <- db
	}()

	db := <-ch
	return db
}
