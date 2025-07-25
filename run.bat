@echo off
echo Starting Tes-Go API Server...
echo.
echo Make sure PostgreSQL is running and the database 'tesgo' exists
echo.
go mod tidy
if %errorlevel% neq 0 (
    echo Failed to download dependencies. Make sure Go is installed.
    pause
    exit /b 1
)

echo Dependencies downloaded successfully.
echo Starting server on port 8080...
echo.
go run main.go
