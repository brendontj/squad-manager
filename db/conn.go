package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func loadDotEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(error(err))
	}
	return os.Getenv(key)
}

// ConnSQL is a function to connect with database
func ConnSQL() *sql.DB {
	host := loadDotEnv("HOST_PG")
	port := loadDotEnv("PORT_PG")
	user := loadDotEnv("USER_PG")
	password := loadDotEnv("PASSWORD_PG")
	dbname := loadDotEnv("DBNAME_PG")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(error(err))
	}
	return db
}
