package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/falschgesagt/Backend/types"
	"github.com/falschgesagt/Backend/utilities"
	"github.com/falschgesagt/Backend/db"
)

func main() {
	db := db.Connection
	defer db.Close()

	appConfig := types.Config{
		IP: utilities.GetStringWithFallbackValueFrom("APP_IP", "localhost"),
		Port: utilities.GetIntWithFallbackValueFrom("APP_PORT", 5432),
	}
	addr := fmt.Sprintf("%s:%s", appConfig.IP, appConfig.Port)
	logOutput := fmt.Sprint("Listening on ", addr)

	http.ListenAndServe(addr, utilities.GetQuotes(db))
	log.Println(logOutput)
}