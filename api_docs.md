 ### 📘 Makerble Backend API Documentation

## 📽️ Demo Video for postman

👉 [Live demo](https://drive.google.com/file/d/1O3_CVSabkrVC-r4H150SfI7zhsL7epRG/view?usp=sharing)

---


## 📽️ Demo Video for curl 

👉 [Live demo](https://drive.google.com/file/d/18rCr5bljQj6BY_TTuo0Ebnb2pwU2MHAz/view?usp=sharing)

---

## 📌 Base URL

```
http://localhost:8080
```

> Use the full URL for each endpoint, like `http://localhost:8080/users`

---

## 🔐 1. Register User

**POST** `http://localhost:8080/users`

### Headers:

```http
Content-Type: application/json
```

### Body:

```json
{
  "name": "Dr. Strange",
  "email": "doctor@makerble.com",
  "password": "secret123",
  "role": "doctor"
}
```

### Response:

```json
{
  "id": 1,
  "name": "Dr. Strange",
  "email": "doctor@makerble.com",
  "role": "doctor"
}
```

### 🔧 curl Example:

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{"name":"Dr. Strange", "email":"doctor@makerble.com", "password":"secret123", "role":"doctor"}'
```

---

## 🔑 2. Login User

**POST** `http://localhost:8080/login`

### Headers:

```http
Content-Type: application/json
```

### Body:

```json
{
  "email": "doctor@makerble.com",
  "password": "secret123"
}
```

### Response:

```json
{
  "message": "Login successful",
  "email": "doctor@makerble.com",
  "role": "doctor"
}
```

### 🔧 curl Example:

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"email":"doctor@makerble.com", "password":"secret123"}'
```

---

## ➕ 3. Create Patient (Receptionist Only)

**POST** `http://localhost:8080/patients`

### Headers:

```http
Content-Type: application/json  
X-User-Email: uma@makerble.com
```

### Body:

```json
{
  "name": "John Doe",
  "age": 30,
  "gender": "Male",
  "diagnosis": "Flu",
  "created_by": 2
}
```

### Response:

```json
{
  "id": 1,
  "name": "John Doe",
  "age": 30,
  "gender": "Male",
  "diagnosis": "Flu",
  "created_by": 2,
  "created_at": "2025-06-24T10:00:00Z"
}
```

### 🔧 curl Example:

```bash
curl -X POST http://localhost:8080/patients \
-H "Content-Type: application/json" \
-H "X-User-Email: uma@makerble.com" \
-d '{"name":"John Doe", "age":30, "gender":"Male", "diagnosis":"Flu", "created_by":2}'
```

---

## 👁 4. Get Patient by ID (Doctor Only)

**GET** `http://localhost:8080/patients/1`

### Headers:

```http
X-User-Email: doctor@makerble.com
```

### Response:

```json
{
  "id": 1,
  "name": "John Doe",
  "age": 30,
  "gender": "Male",
  "diagnosis": "Flu",
  "created_by": 2,
  "created_at": "2025-06-24T10:00:00Z"
}
```

### 🔧 curl Example:

```bash
curl -X GET http://localhost:8080/patients/1 \
-H "X-User-Email: doctor@makerble.com"
```

---

## 🧒 5. Update Patient (Doctor Only)

**PUT** `http://localhost:8080/patients/1`

### Headers:

```http
Content-Type: application/json  
X-User-Email: doctor@makerble.com
```

### Body:

```json
{
  "name": "John Smith",
  "age": 35,
  "gender": "Male",
  "diagnosis": "Covid-19"
}
```

### Response:

```json
{
  "message": "Patient updated successfully"
}
```

### 🔧 curl Example:

```bash
curl -X PUT http://localhost:8080/patients/1 \
-H "Content-Type: application/json" \
-H "X-User-Email: doctor@makerble.com" \
-d '{"name":"John Smith", "age":35, "gender":"Male", "diagnosis":"Covid-19"}'
```

---

## 📋 6. List All Patients (Doctor Only)

**GET** `http://localhost:8080/patients`

### Headers:

```http
X-User-Email: doctor@makerble.com
```

### Response:

```json
[
  {
    "id": 1,
    "name": "John Smith",
    "age": 35,
    "gender": "Male",
    "diagnosis": "Covid-19",
    "created_by": 2,
    "created_at": "2025-06-24T10:00:00Z"
  }
]
```

### 🔧 curl Example:

```bash
curl -X GET http://localhost:8080/patients \
-H "X-User-Email: doctor@makerble.com"
```

---

## ❌ 7. Delete Patient by ID (Doctor Only)

**DELETE** `http://localhost:8080/patients/1`

### Headers:

```http
X-User-Email: doctor@makerble.com
```

### Response:

```json
{
  "message": "Patient deleted successfully"
}
```

### 🔧 curl Example:

```bash
curl -X DELETE http://localhost:8080/patients/1 \
-H "X-User-Email: doctor@makerble.com"
```

---

## 📝 Notes

* Use correct `X-User-Email` to authorize:

  * Doctors can `GET`, `PUT`, `DELETE`
  * Receptionists can `POST` patients
* Gender must be `Male`, `Female`, or `Other` (case-insensitive)
* IDs should match the user/patient created in DB

---

############################### 🔐 POST /login  using curl#######################
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






