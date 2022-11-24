package configs

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db
	// db *gorm.DB
)

func Connect() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=goLANGn1nja")
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
