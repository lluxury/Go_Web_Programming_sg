package data

import (
    "database/sql"
    _ "github.com/lib/pq"
    // "fmt"
    "log"
)

// Var Db *sql.DB
var Db *sql.DB

func init()  {
	var err error
	// Db, err = sql.Open("postgres","dbname=yann" sslmode=disable)
    // Db, 
    Db, err = sql.Open("postgres","postgres:abc@tcp(192.168.142.9:5432)/yann")
	if err != nil{
		log.Fatal(err)
	}

	return 
	
}

