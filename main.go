package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/falschgesagt/Backend/db"
	"github.com/falschgesagt/Backend/types"
	"github.com/falschgesagt/Backend/utilities"
)

func main() {
	db := db.Connection
	defer db.Close()

	appConfig := types.Config{
		IP: utilities.GetStringWithFallbackValueFrom("APP_IP", "localhost"),
		Port: 8080,
	}
	addr := fmt.Sprintf("%s:%d", appConfig.IP, appConfig.Port)
	logOutput := fmt.Sprint("Listening on ", addr)

	log.Println(logOutput)
	http.ListenAndServe(addr, utilities.GetQuotes(db))
}
