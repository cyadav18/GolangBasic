package database

import (
	"fmt"
	_ "fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	dbName   = "NEWTEST"
	port     = "5432"
	sslmode  = "disable"
	TimeZone = "Asia/Calcutta"
)
var db *gorm.DB
func ConnectToDB() {
	//dsn := "host=localhost user=postgres password=password dbname=NEWTEST port=5432 sslmode=disable TimeZone=Asia/Calcutta"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslmode, TimeZone)
	fmt.Println(dsn)
	dbl, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!= nil{
		fmt.Println("DB connection established sucessfully")
	}
	fmt.Println("DB connection established sucessfully")
	db = dbl
}

