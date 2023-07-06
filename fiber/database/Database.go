package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

type Database struct {
	Context *sql.DB
}

func GetConnectionString() string {
	return "Server=174.138.24.220;Database=NotificationService;User Id=sa;password=708@^@r$#@!k#;Trusted_Connection=False;MultipleActiveResultSets=true;"
}

func GetDbContext() *Database {
	var (
		err              error
		db               *sql.DB
		connectionString = GetConnectionString()
	)

	if db, err = sql.Open("mssql", connectionString); err != nil {
		log.Println(fmt.Errorf("error opening database: %v", err))
	}

	return &Database{
		Context: db,
	}
}

func (database *Database) Close() {
	if err := database.Context.Close(); err != nil {
		log.Print(err.Error())
	}
}

func (database *Database) Ping() error {
	var ctx = context.Background()

	if err := database.Context.PingContext(ctx); err != nil {
		return err
	}

	var err error
	var data *sql.Rows

	sqlStatement := fmt.Sprintf("SELECT @@VERSION AS VERSION;")

	if data, err = database.Context.QueryContext(ctx, sqlStatement); err != nil {
		return err
	}

	for data.Next() {
		var version string
		if err := data.Scan(&version); err != nil {
			return err
		}

		log.Println("Version of MSSQL ", version)
	}
	return err
}
