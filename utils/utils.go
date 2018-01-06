package utils

import "os"

func GetPort() string {
	userSpecifiedPort := os.Getenv("PORT")
	port := "5555"

	if userSpecifiedPort != "" {
		port = userSpecifiedPort
	}

	return ":" + port
}
