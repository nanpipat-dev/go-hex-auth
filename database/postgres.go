package database

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Postgres *gorm.DB
}

var dbConnect = &DB{}

func ConnectToPostgreSQL(host, port, username, pass, dbname string) (*DB, error) {
	fmt.Println("User is: ", username)
	fmt.Println("Password is: ", pass)
	fmt.Println("Host is: ", host)
	fmt.Println("Port is: ", port)
	fmt.Println("Dbname is: ", dbname)

	//config string
	var connectionStr string

	if host == "" && port == "" && dbname == "" {
		return nil, errors.New("cannot estabished the connection")
	}

	if port == "APP_DATABASE_POSTGRES_PORT" {
		port = "5432"
	}

	// if sslmode {
	// 	connectionStr = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, username, pass, dbname, port)
	// } else {
	// 	connectionStr = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, username, pass, dbname, port)
	// }

	connectionStr = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, username, pass, dbname, port)

	//connect postgres
	dial := postgres.Open(connectionStr)

	var err error
	pg, err := gorm.Open(dial, &gorm.Config{
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	dbConnect.Postgres = pg
	return dbConnect, nil
}

func DisconnectPostgres(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		panic("close db")
	}
	sqlDb.Close()
	log.Println("Connected with postgres has closed")
}
