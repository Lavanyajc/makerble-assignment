# 📘 Makerble Backend API Documentation

## 🔐 POST /login
Login with email and password.

**Request**
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "email": "sita@makerble.com",
  "password": "lavanya123"
}'
```

**Response**
```json
{
  "role": "receptionist",
  "name": "sitaram Builder",
  "email": "sita@makerble.com"
}
```

---

## 👤 POST /users
Create a new user.

**Request**
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
  "name": "sitaram Builder",
  "email": "sita@makerble.com",
  "password": "lavanya123",
  "role": "receptionist"
}'
```

---

## 🏥 POST /patients
Add a new patient.

**Request**
```bash
curl -X POST http://localhost:8080/patients \
-H "Content-Type: application/json" \
-d '{
  "name": "Ravi Patel",
  "age": 35,
  "gender": "male",
  "diagnosis": "Migraine",
  "created_by": 1
}'
```

---

## 📋 GET /patients
Get list of all patients.

**Request**
```bash
curl http://localhost:8080/patients
```

---

## 🔎 GET /patients/{id}
View a specific patient.

**Request**
```bash
curl http://localhost:8080/patients/1
```

---

## ✏️ PUT /patients/{id}
Update a patient.

**Request**
```bash
curl -X PUT http://localhost:8080/patients/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Ravi P",
  "age": 36,
  "gender": "male",
  "diagnosis": "Updated Diagnosis"
}'
```


