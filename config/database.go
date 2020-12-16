package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// DB var
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig func
func BuildDBConfig() *DBConfig {
	// dbConfig := 
	return &DBConfig{
	Host:     "localhost",
	Port:     3306,
	User:     "root",
	Password: "1234",
	DBName:   "ginTest",
	}
	// return &dbConfig
}

//DbURL func
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
	"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	dbConfig.User,
	dbConfig.Password,
	dbConfig.Host,
	dbConfig.Port,
	dbConfig.DBName,
	)
}
