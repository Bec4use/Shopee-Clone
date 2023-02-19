package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {
	dsn := "host=localhost user=user password=admin dbname=gojwt port=54320 sslmode=disable TimeZone=Asia/Bangkok"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migtrate the schema
	Db.AutoMigrate(&Users{})
}
