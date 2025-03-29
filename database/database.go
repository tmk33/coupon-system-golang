package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=brolf1dmcfzalsouabk7-postgresql.services.clever-cloud.com user=uuesdky9jyzk5m1odjf9 password=gMSQVvn4dgMwalvtY2YUPbaPJnZVdL dbname=brolf1dmcfzalsouabk7 port=50013 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db
}
