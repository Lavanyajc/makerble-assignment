package config

import (
	"fmt"
)

// CreateUsersTable auto-creates the 'users' table in PostgreSQL
func CreateUsersTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
                role TEXT CHECK (role IN ('doctor', 'receptionist')) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(query)
	if err != nil {
		panic("❌ Failed to create users table: " + err.Error())
	}

	fmt.Println("✅ Users table created (or already exists).")
}
// CreateVisitLogsTable auto-creates the 'visit_logs' table in PostgreSQL
func CreateVisitLogsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS visit_logs (
		id SERIAL PRIMARY KEY,
		user_id INT,
		endpoint TEXT NOT NULL,
		method TEXT NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
	);`

	_, err := DB.Exec(query)
	if err != nil {
		panic("❌ Failed to create visit_logs table: " + err.Error())
	}

	fmt.Println("✅ VisitLogs table created (or already exists).")
}
