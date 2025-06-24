
---

## 📁 1. Navigate to your project

```bash
cd ~/Downloads/makerble-clean
```

---

## ⚙️ 2. Initialize Go project (only once)

```bash
go mod init makerble-clean
```

---

## 📦 3. Install Dependencies

```bash
go get github.com/lib/pq
go get github.com/joho/godotenv
go get github.com/gorilla/mux
```

---

## 🚀 4. Run Your Application

```bash
go run main.go
```

You should see:

```
✅ Successfully connected to PostgreSQL!
✅ Users table created (or already exists).
✅ VisitLogs table created (or already exists).
🚀 Server is running on port 8080...
```

---

## 🧪 5. CURL Commands to Test APIs

---

### 👤 Create a User (POST /users)

```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
  "name": "Lavanya Dev",
  "email": "lavanya@makerble.com",
  "password": "lavanya123",
  "role": "receptionist"
}'
```

---

### 🏥 Create a Patient (POST /patients)

```bash
curl -X POST http://localhost:8080/patients \
-H "Content-Type: application/json" \
-d '{
  "name": "Ravi Patel",
  "age": 30,
  "gender": "male",
  "diagnosis": "Flu",
  "created_by": 1
}'
```

---

### 📋 Get All Patients (GET /patients)

```bash
curl http://localhost:8080/patients
```

---

### 🔍 Get a Specific Patient (GET /patients/{id})

```bash
curl http://localhost:8080/patients/1
```

---

### ✏️ Update a Patient (PUT /patients/{id})

```bash
curl -X PUT http://localhost:8080/patients/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Ravi P",
  "age": 31,
  "gender": "male",
  "diagnosis": "Recovered"
}'
```

---

## 🔒 Sample Login (if implemented)

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "email": "lavanya@makerble.com",
  "password": "lavanya123"
}'
```

---

## 📄 6. Create Important Files (if not done yet)

```bash
touch .gitignore
echo ".env" >> .gitignore

touch api_docs.md
touch README.md
```

---

## 🔁 7. Git Commands for Push

```bash
git init
git add .
git commit -m "✅ Makerble backend assignment complete"
git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
git branch -M main
git push -u origin main
```

---




