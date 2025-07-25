# Test API Endpoints

## 1. Login API Test

### Request
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "card_number": "1234567890",
    "password": "password123"
  }'
```

### Expected Response
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "card_number": "1234567890", 
    "name": "John Doe",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

## 2. Create Terminal API Test

### Request (Replace YOUR_JWT_TOKEN with actual token from login)
```bash
curl -X POST http://localhost:8080/terminals \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "name": "Terminal Utama",
    "code": "TRM001", 
    "location": "Jakarta Pusat"
  }'
```

### Expected Response
```json
{
  "message": "Terminal created successfully",
  "terminal": {
    "id": 1,
    "name": "Terminal Utama",
    "code": "TRM001",
    "location": "Jakarta Pusat",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

## 3. Error Cases

### Login with Invalid Credentials
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "card_number": "invalid",
    "password": "wrong"
  }'
```

### Create Terminal without Authorization
```bash
curl -X POST http://localhost:8080/terminals \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Terminal Test",
    "code": "TRM002",
    "location": "Test Location"
  }'
```

### Create Terminal with Invalid Token
```bash
curl -X POST http://localhost:8080/terminals \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer invalid_token" \
  -d '{
    "name": "Terminal Test",
    "code": "TRM002", 
    "location": "Test Location"
  }'
```
