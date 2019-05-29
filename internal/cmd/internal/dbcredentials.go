// +build db

package internal

import "os"

func GetDBCredentialsEnvVar() (dbHost, dbPort, dbUser, dbPass string) {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	return
}
