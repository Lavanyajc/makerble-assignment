package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"makerble-clean/config"
	"makerble-clean/handlers"
)

func main() {
	config.ConnectDB()
	config.CreateUsersTable()
	config.CreateVisitLogsTable()

	r := mux.NewRouter().StrictSlash(true)

	// ✅ User APIs
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// ✅ Patient APIs
	r.HandleFunc("/patients", handlers.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", handlers.GetAllPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", handlers.GetPatientByID).Methods("GET")
	r.HandleFunc("/patients/{id}", handlers.DeletePatient).Methods("DELETE")

	// ✅ Root route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hii Lavanya , API is live!"))
	})

	// ✅ Server start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	fmt.Printf("🚀 Server is running on port %s...\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
