package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
        "github.com/gorilla/mux"
	"makerble-clean/config"
	"makerble-clean/models"
)

// POST /patients
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Println("❌ Error decoding patient:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO patients (name, age, gender, diagnosis, created_by)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	err = config.DB.QueryRow(query,
		patient.Name,
		patient.Age,
		patient.Gender,
		patient.Diagnosis,
		patient.CreatedBy,
	).Scan(&patient.ID, &patient.CreatedAt)

	if err != nil {
		log.Println("❌ DB Error:", err)
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}

	config.LogVisit("/patients", "POST")

	log.Println("✅ Patient created:", patient.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

// GET /patients
func GetAllPatients(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT id, name, age, gender, diagnosis, created_by, created_at FROM patients`)
	if err != nil {
		log.Println("❌ DB Error:", err)
		http.Error(w, "Failed to fetch patients", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var patients []models.Patient

	for rows.Next() {
		var p models.Patient
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Gender, &p.Diagnosis, &p.CreatedBy, &p.CreatedAt)
		if err != nil {
			log.Println("❌ Scan Error:", err)
			continue
		}
		patients = append(patients, p)
	}

	config.LogVisit("/patients", "GET")

	log.Println("✅ Patients fetched:", len(patients))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}


func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var updated models.Patient
    err := json.NewDecoder(r.Body).Decode(&updated)
    if err != nil {
        log.Println("❌ Error decoding update:", err)
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // ✅ Extract email from header
    email := r.Header.Get("X-User-Email")
    if email == "" {
        http.Error(w, "Missing X-User-Email header", http.StatusUnauthorized)
        return
    }

    // ✅ Fetch role from DB
    var role string
    err = config.DB.QueryRow("SELECT role FROM users WHERE email = $1", email).Scan(&role)
    if err != nil {
        log.Println("❌ Failed to get user role:", err)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // ✅ Check if user is allowed to update
    if role != "doctor" && role != "receptionist" {
        http.Error(w, "Only doctors or receptionists can update patients", http.StatusForbidden)
        return
    }

    // 🛠️ Update patient
    query := `UPDATE patients SET name=$1, age=$2, gender=$3, diagnosis=$4 WHERE id=$5 RETURNING created_by, created_at`
    err = config.DB.QueryRow(query, updated.Name, updated.Age, updated.Gender, updated.Diagnosis, id).
        Scan(&updated.CreatedBy, &updated.CreatedAt)
    if err != nil {
        log.Println("❌ DB Update Error:", err)
        http.Error(w, "Update failed", http.StatusInternalServerError)
        return
    }

    updated.ID = parseID(id)
    config.LogVisit("/patients/"+id, "PUT")

    log.Println("✅ Patient updated:", id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updated)
}


func parseID(s string) int {
	var id int
	fmt.Sscanf(s, "%d", &id)
	return id
}

func GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var patient models.Patient
	query := `SELECT id, name, age, gender, diagnosis, created_by, created_at FROM patients WHERE id=$1`

	err := config.DB.QueryRow(query, id).Scan(
		&patient.ID,
		&patient.Name,
		&patient.Age,
		&patient.Gender,
		&patient.Diagnosis,
		&patient.CreatedBy,
		&patient.CreatedAt,
	)

	if err != nil {
		log.Println("❌ Patient fetch error:", err)
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	config.LogVisit("/patients/"+id, "GET")

	log.Println("✅ Patient fetched:", patient.Name)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patient)
}

