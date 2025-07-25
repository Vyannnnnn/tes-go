# Setup Instructions

## Prerequisites

Before running this API, you need to install:

1. **Go Programming Language**
   - Download from: https://golang.org/dl/
   - Choose the Windows installer for your system (64-bit recommended)
   - Follow the installation wizard
   - Verify installation by opening Command Prompt and running: `go version`

2. **PostgreSQL Database**
   - Download from: https://www.postgresql.org/download/windows/
   - Install PostgreSQL with default settings
   - Remember the password you set for the 'postgres' user
   - Create a database named 'tesgo':
     ```sql
     CREATE DATABASE tesgo;
     ```

## Installation and Running

1. **Open Command Prompt or PowerShell**
   - Navigate to the project directory: `cd "c:\Users\user\Downloads\Tes-Go"`

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Configure Database Connection (Optional)**
   - Edit `main.go` file if your PostgreSQL setup is different
   - Default connection: `user=postgres dbname=tesgo sslmode=disable password=password host=localhost`

4. **Run the API**
   ```bash
   go run main.go
   ```
   
   Or double-click `run.bat` file

5. **Test the API**
   - The server will start on `http://localhost:8080`
   - Use the test commands in `test_endpoints.md`

## Quick Test

Once the server is running, test the login endpoint:

```bash
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"card_number\":\"1234567890\",\"password\":\"password123\"}"
```

## Troubleshooting

- **"go command not found"**: Install Go from https://golang.org/dl/
- **Database connection error**: Make sure PostgreSQL is running and the database 'tesgo' exists
- **Port 8080 already in use**: Change the port in the code or set environment variable `PORT=8081`
