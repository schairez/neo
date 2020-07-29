package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	d, err := sql.Open("postgres", getDBConnURI())
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(db.NewDB(d), cors)
	err = app.Serve()
	log.Println("Error", err)
}

//driver specific data source name
const (
	dialect  = "postgresql"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "neo_demo"
)

func getDBConnURI() string {
	// host := "localhost"
	// pass := "pass"

	// if os.Getenv("profile") == "prod" {
	// 	host = "db"
	// 	pass = os.Getenv("db_pass")
	// }
	// fmt.Sprintf()
	fmt.Sprintf("%s://%v")
	return dialect + "://" + host + ":5432"
	return "postgresql://" + host + ":5432/pos" +
		"?user=goxygen&sslmode=disable&password=" + pass

}

/*
serverName=localhost
databaseName=test
user=testuser
password=testpassword

postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
dialect+driver://username:password@host:port/database

postgresql://user:pass@localhost:5432/my_db

export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'


DATA SOURCE NAME
$ export DSN=postgres://hydra:secret@ory-hydra-example--postgres:5432/hydra?sslmode=disable



*/
