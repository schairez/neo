package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres golang driver
)

// const (
// 	// TODO fill this in directly or through environment variable
// 	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
// 	dbDSN = ""
// )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "neo_demo"
	sslmode  = "disable"
)

//export POSTGRESQL_URL='postgres://postgres:password@localhost:5432/example?sslmode=disable'


//


// var db *sqlx.DB

//PostgresDB struct
type PostgresDB struct {
	db *sqlx.DB
}

// type DB interface {
// 	GetTodos() ([]*)
// }
