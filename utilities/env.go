package utilities

import (
	"log"
	"os"
	"strconv"
)

func GetStringWithFallbackValueFrom(key, fallbackValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return fallbackValue
	}
	return v
}

func GetStringFrom(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalln("could not find " + key + " on environment variables")
	}
	return v
}

func GetIntWithFallbackValueFrom(key string, fallbackValue int) int {
	v := GetStringWithFallbackValueFrom(key, "")

	r, err := strconv.Atoi(v)
	if err != nil || v == "" {
		return fallbackValue
	}
	return r
}