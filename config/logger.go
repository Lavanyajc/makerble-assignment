package config

import "log"

// LogVisit inserts a record into visit_logs
func LogVisit(endpoint, method string) {
	query := `
        INSERT INTO visit_logs (endpoint, method)
        VALUES ($1, $2);
    `

	_, err := DB.Exec(query, endpoint, method)
	if err != nil {
		log.Println("❌ Failed to log visit:", err)
		return
	}

	log.Println("📌 Visit logged:", method, endpoint)
}

