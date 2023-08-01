package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

func FromSqlRaw[Result comparable](db *Database, sqlStatement string, params func(result *Result) []any) (*[]Result, error) {
	var (
		err     error
		results []Result
		data    *sql.Rows
		ctx     = context.Background()
	)

	if data, err = db.Context.QueryContext(ctx, sqlStatement); err != nil {
		log.Println(err.Error())
		return &results, err
	}

	for data.Next() {
		var result Result
		var dest = params(&result)

		if err = data.Scan(dest...); err != nil {
			log.Println(err.Error())
			return &results, err
		}

		results = append(results, result)
	}

	return &results, err
}

func FromStoredProcedure[Result comparable](db *Database, query string, args []any, params func(result *Result) []any) (*[]Result, error) {
	var (
		err     error
		results []Result
		data    *sql.Rows
		ctx     = context.Background()
	)

	data, err = db.Context.QueryContext(ctx, fmt.Sprintf("EXEC %s", query), args...)

	for data.Next() {
		var result Result
		var dest = params(&result)

		if err = data.Scan(dest...); err != nil {
			log.Println(err.Error())
			return &results, err
		}

		results = append(results, result)
	}

	return &results, err
}
