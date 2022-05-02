package database

import (
	"goserver/app/config"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}
type Database struct {
	*gorm.DB
}

var db *gorm.DB

func ConnectDB(dbconfig *config.Config) (*Database, error) {
	var err error
	config := DatabaseConfig{
		Driver:   dbconfig.GetString("DB_DRIVER"),
		Host:     dbconfig.GetString("DB_HOST"),
		Username: dbconfig.GetString("DB_USERNAME"),
		Password: dbconfig.GetString("DB_PASSWORD"),
		Port:     dbconfig.GetInt("DB_PORT"),
		Database: dbconfig.GetString("DB_DATABASE"),
	}
	//MySQL Connection format
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return &Database{db}, err
}

func GetDB() *gorm.DB {
	return db
}
