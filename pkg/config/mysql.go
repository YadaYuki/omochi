package config

import (
	"fmt"
	"os"
)

type MysqlConfigType struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
}

var MysqlConfig = MysqlConfigType{
	DbUser:     os.Getenv("MYSQL_USER"),
	DbPassword: os.Getenv("MYSQL_PASSWORD"),
	DbHost:     os.Getenv("MYSQL_HOST"),
	DbName:     os.Getenv("MYSQL_DATABASE"),
	DbPort:     os.Getenv("DB_PORT"),
}

var MysqlConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", MysqlConfig.DbUser, MysqlConfig.DbPassword, MysqlConfig.DbHost, MysqlConfig.DbPort, MysqlConfig.DbName)
