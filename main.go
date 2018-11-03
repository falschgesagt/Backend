package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/falschgesagt/Backend/types"
	"github.com/falschgesagt/Backend/utilities"
	"log"
	"net/http"

	"github.com/falschgesagt/Backend/db"
)

// Output Output struct used for outputting right JSON object
type Output struct {
	Quotes  []string `json:"quotes"`
	Authors []string `json:"authors"`
}

// GetQuotes Gets all Qutoes
func GetQuotes(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		quotes := make([]string, 0)
		authors := make([]string, 0)

		quotesQuery := utilities.QueryAndPanicOnError("SELECT * FROM quote ORDER BY RAND();", db)
		defer quotesQuery.Close()

		authorsQuery := utilities.QueryAndPanicOnError("SELECT * FROM author ORDER BY RAND() LIMIT 500;", db)
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

		var output = Output{quotes, authors}
		js, err := json.Marshal(output)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func main() {
	db := db.Connection
	defer db.Close()

	appConfig := types.Config{
		IP: utilities.GetStringWithFallbackValueFrom("APP_IP", "localhost"),
		Port: utilities.GetIntWithFallbackValueFrom("APP_PORT", 5432),
	}
	addr := fmt.Sprintf("%s:%s", appConfig.IP, appConfig.Port)
	logOutput := fmt.Sprint("Listening on ", addr)

	http.ListenAndServe(addr, GetQuotes(db))
	log.Println(logOutput)
}