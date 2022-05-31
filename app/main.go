package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	repository "github.com/AndyMile/candyStore/repository"
	router "github.com/AndyMile/candyStore/provider/rest"
	controller "github.com/AndyMile/candyStore/controllers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := DBconnection(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()

	storeRepo := repository.NewStoreRepo(db)
	c := controller.NewBaseHandler(storeRepo)
	router.RouterHandler(c)
}

func DBconnection(user string, password string, host string, dbname string) *gorm.DB {
	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}