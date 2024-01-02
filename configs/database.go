package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func databaseError(err error) {
	if err != nil {
		log.Fatalf("Failed connect database: %v", err)
	}
}

func NewDatabase(viper *viper.Viper) (db *sql.DB) {
	host := viper.GetString("DB_HOST")
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	port := viper.GetString("DB_PORT")
	database := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	databaseError(err)

	err = db.Ping()
	databaseError(err)

	fmt.Println("Database Connected")
	return db
}
