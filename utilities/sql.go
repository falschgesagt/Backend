package utilities

import "database/sql"

// ExecuteAndPanicOnError Executes a given query and panics if error occurs.
func ExecuteAndPanicOnError(query string, db *sql.DB) {
	_, err := db.Exec(query)

	if err != nil {
		panic(err)
	}
}


// QueryAndPanicOnError Executes a given query, panic if error occurs and if successfull return the rows
func QueryAndPanicOnError(query string, db *sql.DB) *sql.Rows {
	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	return rows
}