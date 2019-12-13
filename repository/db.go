package repository

import (
	"fmt"

	"github.com/anaabdi/todo-go/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func DB() *sqlx.DB {
	return db
}

func InitDB() {
	connStr := fmt.Sprintf(
		"dbname=%s client_encoding=UTF8 user=%s password=%s host=%s port=%s sslmode=%s",
		config.GetString("DB_NAME", "todo"),
		config.GetString("DB_USER", "postgres"),
		config.GetString("DB_PASSWORD", "1234"),
		config.GetString("DB_HOST", "localhost"),
		config.GetString("DB_PORT", "5432"),
		config.GetString("DB_SSL_MODE", "disable"),
	)

	db = sqlx.MustConnect("postgres", connStr)
}

var (
	queryInsert = `INSERT INTO users (username, profile, created_by, updated_by) 
	VALUES($1, $2, $3, $4) RETURNING id`

	queryUpdateUserID = `UPDATE users SET profile = $1 WHERE id = $2`

	queryUserByUsername = `SELECT id, username, profile, deactivated_at FROM users 
	WHERE deleted_at IS NULL AND username = $1`
)
