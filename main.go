package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Database connection
var db *sql.DB

// JWT secret key - in production, this should be an environment variable
var jwtSecret = []byte("your-secret-key")

// User struct
type User struct {
	ID         int       `json:"id" db:"id"`
	CardNumber string    `json:"card_number" db:"card_number"`
	Name       string    `json:"name" db:"name"`
	Password   string    `json:"password,omitempty" db:"password"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// Terminal struct
type Terminal struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Code      string    `json:"code" db:"code"`
	Location  string    `json:"location" db:"location"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Login request struct
type LoginRequest struct {
	CardNumber string `json:"card_number" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// Login response struct
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    User   `json:"user"`
}

// Terminal request struct
type TerminalRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Location string `json:"location" binding:"required"`
}

// JWT Claims
type Claims struct {
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	jwt.RegisteredClaims
}

// Initialize database connection
func initDB() {
	var err error
	
	// Database connection string - adjust according to your PostgreSQL setup
	connStr := "user=postgres dbname=tesgo sslmode=disable password=password host=localhost"
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		connStr = dbURL
	}
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to database successfully")
	
	// Create tables if they don't exist
	createTables()
}

// Create database tables
func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			card_number VARCHAR(255) UNIQUE NOT NULL,
			name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS terminals (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			code VARCHAR(255) UNIQUE NOT NULL,
			location TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS prepaid_cards (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id),
			balance DECIMAL(10,2) DEFAULT 0,
			last_sync_at TIMESTAMP,
			status VARCHAR(50) DEFAULT 'active',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS gates (
			id SERIAL PRIMARY KEY,
			terminal_id INTEGER NOT NULL REFERENCES terminals(id),
			gate_code VARCHAR(255) NOT NULL,
			is_active BOOLEAN DEFAULT true,
			last_online TIMESTAMP,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			card_id INTEGER NOT NULL REFERENCES prepaid_cards(id),
			user_id INTEGER NOT NULL REFERENCES users(id),
			checkin_terminal_id INTEGER REFERENCES terminals(id),
			checkout_terminal_id INTEGER REFERENCES terminals(id),
			checkin_time TIMESTAMP,
			checkout_time TIMESTAMP,
			fare DECIMAL(10,2),
			status VARCHAR(50),
			sync_status VARCHAR(50),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS fares (
			id SERIAL PRIMARY KEY,
			from_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
			to_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
			fare_amount DECIMAL(10,2) NOT NULL,
			effective_from DATE NOT NULL,
			effective_to DATE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS transaction_logs (
			id SERIAL PRIMARY KEY,
			gate_id INTEGER NOT NULL REFERENCES gates(id),
			card_id INTEGER NOT NULL REFERENCES prepaid_cards(id),
			log_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			event_type VARCHAR(100) NOT NULL,
			raw_data TEXT,
			is_synced BOOLEAN DEFAULT false,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Printf("Error creating table: %v", err)
		}
	}

	// Insert sample user if not exists
	insertSampleData()
}

// Insert sample data for testing
func insertSampleData() {
	// Check if sample user exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE card_number = $1", "1234567890").Scan(&count)
	if err != nil {
		log.Printf("Error checking sample user: %v", err)
		return
	}

	if count == 0 {
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %v", err)
			return
		}

		// Insert sample user
		_, err = db.Exec(`
			INSERT INTO users (card_number, name, password) 
			VALUES ($1, $2, $3)`,
			"1234567890", "John Doe", string(hashedPassword))
		if err != nil {
			log.Printf("Error inserting sample user: %v", err)
		} else {
			log.Println("Sample user created: card_number=1234567890, password=password123")
		}
	}
}

// Generate JWT token
func generateJWT(userID int, cardNumber string) (string, error) {
	claims := Claims{
		UserID:     userID,
		CardNumber: cardNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Middleware to validate JWT token
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix
		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserID)
		c.Set("card_number", claims.CardNumber)
		c.Next()
	}
}

// Login API handler
func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by card number
	var user User
	var hashedPassword string
	err := db.QueryRow(`
		SELECT id, card_number, name, password, created_at, updated_at 
		FROM users WHERE card_number = $1`,
		req.CardNumber).Scan(
		&user.ID, &user.CardNumber, &user.Name, &hashedPassword,
		&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid card number or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid card number or password"})
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.ID, user.CardNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Remove password from response
	user.Password = ""

	response := LoginResponse{
		Message: "Login successful",
		Token:   token,
		User:    user,
	}

	c.JSON(http.StatusOK, response)
}

// Create Terminal API handler
func createTerminal(c *gin.Context) {
	var req TerminalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user info from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Check if terminal code already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM terminals WHERE code = $1", req.Code).Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Terminal code already exists"})
		return
	}

	// Insert new terminal
	var terminal Terminal
	err = db.QueryRow(`
		INSERT INTO terminals (name, code, location) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, code, location, created_at`,
		req.Name, req.Code, req.Location).Scan(
		&terminal.ID, &terminal.Name, &terminal.Code, &terminal.Location, &terminal.CreatedAt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create terminal"})
		return
	}

	log.Printf("Terminal created by user ID %v: %+v", userID, terminal)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Terminal created successfully",
		"terminal": terminal,
	})
}

// Get all terminals (bonus endpoint)
func getTerminals(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, code, location, created_at FROM terminals ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var terminals []Terminal
	for rows.Next() {
		var terminal Terminal
		err := rows.Scan(&terminal.ID, &terminal.Name, &terminal.Code, &terminal.Location, &terminal.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning terminals"})
			return
		}
		terminals = append(terminals, terminal)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Terminals retrieved successfully",
		"terminals": terminals,
	})
}

// Health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"message":   "API is running",
	})
}

func main() {
	// Initialize database
	initDB()
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Public routes
	router.GET("/health", healthCheck)
	router.POST("/login", login)

	// Protected routes (require JWT authentication)
	auth := router.Group("/")
	auth.Use(authMiddleware())
	{
		auth.POST("/terminals", createTerminal)
		auth.GET("/terminals", getTerminals)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Printf("Available endpoints:")
	log.Printf("  POST /login - User login")
	log.Printf("  POST /terminals - Create terminal (requires auth)")
	log.Printf("  GET /terminals - Get all terminals (requires auth)")
	log.Printf("  GET /health - Health check")

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
