package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"makerble-clean/config"
	"makerble-clean/models"
)

// Util function to write JSON response
func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("❌ Error decoding JSON:", err)
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"status":  "fail",
			"message": "Invalid input JSON",
		})
		return
	}

	// Insert user
	sqlStatement := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	err = config.DB.QueryRow(sqlStatement, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		log.Println("❌ DB insert error:", err)
		writeJSON(w, http.StatusConflict, map[string]string{
			"status":  "fail",
			"message": "Email already exists",
		})
		return
	}

	// Log visit
	config.LogVisit("/users", "POST")
	log.Println("✅ New user created:", user.Email)

	// Respond
	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}

// POST /login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"status":  "fail",
			"message": "Invalid login input",
		})
		return
	}

	// Fetch user from DB
	var user models.User
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE email=$1`
	err = config.DB.QueryRow(query, creds.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt,
	)

	// Debug logs
	log.Println("🔍 DB Email:", user.Email)
	log.Println("🔐 DB Password:", user.Password)
	log.Println("🆚 Input Password:", creds.Password)

	if err != nil || user.Password != creds.Password {
		log.Println("❌ Login failed for:", creds.Email)
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"status":  "fail",
			"message": "Invalid email or password",
		})
		return
	}

	// Log visit
	config.LogVisit("/login", "POST")
	log.Println("✅ Login successful for:", user.Email)

	// Respond
	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": map[string]string{
			"role":  user.Role,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

