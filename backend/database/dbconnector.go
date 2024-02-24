package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// DBConnection represents the database connection.
type DBConnection struct {
	db *sql.DB
}

// NewDBConnection creates a new database connection.
func NewDBConnection() (*DBConnection, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DBConnection{db: db}, nil
}

// Close closes the database connection.
func (conn *DBConnection) Close() {
	if conn.db != nil {
		conn.db.Close()
	}
}

// ExecuteQuery executes the provided SQL query template with the given arguments and returns the result.
func (conn *DBConnection) ExecuteQuery(queryTemplate string, args ...interface{}) (*sql.Rows, error) {
	rows, err := conn.db.Query(queryTemplate, args...)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	return rows, nil
}
