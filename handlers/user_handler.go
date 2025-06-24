package handlers

import (
	
	"encoding/json"
	"log"
	"net/http"

	"makerble-clean/config"
	"makerble-clean/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse JSON from request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("❌ Error decoding JSON:", err)
		http.Error(w, `{"status":"fail","message":"Invalid input"}`, http.StatusBadRequest)
		return
	}

	// Insert into users table
	sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	err = config.DB.QueryRow(sqlStatement, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		log.Println("❌ DB error:", err)
		http.Error(w, `{"status":"fail","message":"Email already exists or DB issue"}`, http.StatusBadRequest)
		return
	}

	// ✅ Log visit to visit_logs table
	config.LogVisit("/users", "POST")

	// ✅ Terminal log
	log.Println("✅ New user created:", user.Email)

	// ✅ Custom JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"status": "success",
		"data":   user,
	}

	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE email=$1`

	err = config.DB.QueryRow(query, creds.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt,
	)

	// ✅ Add logs HERE ⬇️
	log.Println("🔎 DB Email:", user.Email)
	log.Println("🔐 DB Password:", user.Password)
	log.Println("🆚 Input Password:", creds.Password)

	if err != nil || user.Password != creds.Password {
		log.Println("❌ Unauthorized login attempt for:", creds.Email)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	config.LogVisit("/login", "POST")
	log.Println("✅ Login successful for:", user.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"role":  user.Role,
		"name":  user.Name,
		"email": user.Email,
	})
}


