package database

import (
	"context"
	"database/sql"
	"log"
)

func FromSqlRaw[Result comparable](db *Database, sqlStatement string, params func(result *Result) []any) *[]Result {
	var (
		err     error
		results []Result
		data    *sql.Rows
		ctx     = context.Background()
	)

	if data, err = db.Context.QueryContext(ctx, sqlStatement); err != nil {
		log.Println(err.Error())
		return &results
	}

	for data.Next() {
		var result Result
		var dest = params(&result)

		if err = data.Scan(dest...); err != nil {
			log.Println(err.Error())
			return &results
		}

		results = append(results, result)
	}

	return &results
}
