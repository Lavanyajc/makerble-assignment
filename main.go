package main

import (
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("🚀 Server is running on port 8080...")
	port := os.Getenv("PORT")
       if port == "" {
         port = "8080" // fallback for local dev
          }
           fmt.Printf("🚀 Server is running on port %s...\n", port)
           log.Fatal(http.ListenAndServe(":"+port, r))

}



