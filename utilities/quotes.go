package utilities

import (
	"database/sql"
	"encoding/json"
	"github.com/falschgesagt/Backend/types"
	"log"
	"net/http"
)

// GetQuotes Gets all Qutoes
func GetQuotes(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quotes := make([]string, 0)
		authors := make([]string, 0)

		quotesQuery := utilities.QueryAndPanicOnError("SELECT * FROM quote ORDER BY RAND();", db)
		defer quotesQuery.Close()

		authorsQuery := utilities.QueryAndPanicOnError("SELECT * FROM author ORDER BY RAND();", db)
		defer authorsQuery.Close()

		for quotesQuery.Next() {
			var quote string

			if err := quotesQuery.Scan(&quote); err != nil {
				log.Fatal(err)
			}

			quotes = append(quotes, quote)
		}

		for authorsQuery.Next() {
			var author string

			if err := authorsQuery.Scan(&author); err != nil {
				log.Fatal(err)
			}

			authors = append(authors, author)
		}

		var output = types.Output{quotes, authors}
		js, err := json.Marshal(output)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}