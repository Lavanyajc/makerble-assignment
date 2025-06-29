## 🏥 Makerble Backend Assignment — Golang + PostgreSQL
---

## 📽️ Demo Video (optional)

👉 [Live demo](https://drive.google.com/file/d/18rCr5bljQj6BY_TTuo0Ebnb2pwU2MHAz/view?usp=sharing)

---


Deployed link - https://makerble-backend-h7fr.onrender.com/


## 🔥 Description

A backend web application built with Golang and PostgreSQL that simulates a hospital system with two portals:
- 👩‍💼 Receptionist Portal (Can register and manage patients)
- 🧑‍⚕️ Doctor Portal (Can view & update patient details)

---

## 🚀 Features

✅ Login API (for both doctors and receptionists)  
✅ Patient CRUD operations  
✅ Role-based data access  
✅ Visit logging for each endpoint  
✅ PostgreSQL database with pgAdmin  
✅ curl-based testing (no Postman used)  
✅ Modular Go codebase with proper folder structure

---

## 💻 Tech Stack

- Golang (`net/http`, `mux`)
- PostgreSQL
- pgAdmin
- curl
- Git Bash

---

## 📦 Project Structure

```
makerble-clean/
├── config/
│   └── db.go
├── handlers/
│   ├── user_handler.go
│   └── patient_handler.go
├── models/
│   ├── user.go
│   └── patient.go
├── .env
├── main.go
├── README.md
└── api_docs.md
```

---

## 🧪 How to Run

1. Ensure PostgreSQL is installed and running
2. Create a database `makerble_db`
3. Set credentials in `.env`:

```
DB_HOST=localhost  
DB_PORT=5432  
DB_USER=postgres  
DB_PASSWORD=yourpassword  
DB_NAME=makerble_db  
```

4. Run the app:

```bash
go run main.go
```

Server will start at: `http://localhost:8080`

---

## 📘 API Documentation

See [`api_docs.md`](./api_docs.md) for full list of endpoints and curl requests.



## 📤 Submission

Submitted as part of the Makerble Golang Assignment by Lavanya JC.


