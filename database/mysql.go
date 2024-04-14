package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type MysqlDatabase struct {
	db *gorm.DB
}

func Initial() *gorm.DB {
	env, _ := godotenv.Read(".env")
	dbURL := env["DATABASE_URL"]
	dbUser := env["DATABASE_USER"]
	dbPwd := env["DATABASE_PWD"]
	dbName := env["DATABASE_NAME"]
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbURL, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error when initial database. %v", err)
	}

	return db
}
