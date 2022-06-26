package env

import (
	"os"
)

type Config struct {
}

var DB_USER = os.Getenv("MYSQL_USER")
var DB_PASSWORD = os.Getenv("MYSQL_PASSWORD")
var DB_HOST = os.Getenv("MYSQL_HOST")
var DB_NAME = os.Getenv("MYSQL_DATABASE")
var DB_PORT = os.Getenv("DB_PORT")
