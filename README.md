# Tes-Go API

A simple Go API for a prepaid card system with JWT authentication.

## Features

- JWT-based authentication
- User login with card number and password
- Terminal management with authorization
- PostgreSQL database integration
- RESTful API design

## Database Schema

The API uses the database schema defined in the DBML format provided, including:
- Users table with card numbers and authentication
- Terminals table for managing transit terminals
- Prepaid cards, transactions, fares, gates, and transaction logs

## API Endpoints

### Public Endpoints

#### POST /login
Login with card number and password to receive a JWT token.

**Request:**
```json
{
  "card_number": "1234567890",
  "password": "password123"
}
```

**Response:**
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

#### GET /health
Health check endpoint.

**Response:**
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "message": "API is running"
}
```

### Protected Endpoints (Require JWT Token)

#### POST /terminals
Create a new terminal. Requires Authorization header with JWT token.

**Headers:**
```
Authorization: Bearer <jwt_token>
```

**Request:**
```json
{
  "name": "Terminal Utama",
  "code": "TRM001",
  "location": "Jakarta Pusat"
}
```

**Response:**
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

#### GET /terminals
Get all terminals. Requires Authorization header with JWT token.

**Headers:**
```
Authorization: Bearer <jwt_token>
```

**Response:**
```json
{
  "message": "Terminals retrieved successfully",
  "terminals": [
    {
      "id": 1,
      "name": "Terminal Utama",
      "code": "TRM001",
      "location": "Jakarta Pusat",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

## Setup and Installation

### Prerequisites

- Go 1.21 or higher
- PostgreSQL database

### Installation Steps

1. Clone or download the project files
2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up PostgreSQL database:
   - Create a database named `tesgo`
   - Update the database connection string in `main.go` if needed

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on port 8080 by default.

### Environment Variables

- `DATABASE_URL`: PostgreSQL connection string (optional, defaults to local setup)
- `PORT`: Server port (optional, defaults to 8080)

### Sample Data

The application automatically creates a sample user for testing:
- Card Number: `1234567890`
- Password: `password123`

## Testing

You can test the API using curl or any HTTP client:

1. **Login:**
   ```bash
   curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"card_number":"1234567890","password":"password123"}'
   ```

2. **Create Terminal (use the token from login response):**
   ```bash
   curl -X POST http://localhost:8080/terminals \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer YOUR_JWT_TOKEN" \
     -d '{"name":"Terminal Utama","code":"TRM001","location":"Jakarta Pusat"}'
   ```

3. **Get Terminals:**
   ```bash
   curl -X GET http://localhost:8080/terminals \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
   ```

## Security

- Passwords are hashed using bcrypt
- JWT tokens expire after 24 hours
- API includes CORS headers for web client support
- Authorization is required for protected endpoints

## Database Schema

The application automatically creates all necessary tables based on the DBML schema provided, including proper foreign key relationships and constraints.
