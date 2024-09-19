package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/models"
	"backend/routes"
	"backend/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}
	return nil
}

func getCredentialsFromEnv() *Credential {
	return &Credential{
		Host:         os.Getenv("DB_HOST"),
		Username:     os.Getenv("DB_USERNAME"),
		Password:     os.Getenv("DB_PASSWORD"),
		DatabaseName: os.Getenv("DB_NAME"),
		Port:         mustGetEnvAsInt("DB_PORT"),
	}
}

func mustGetEnvAsInt(key string) int {
	value := os.Getenv(key)
	port, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid value for %s: %v", key, err)
	}
	return port
}

func Connect(creds *Credential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the dbConn in the context
		c.Set("dbConn", db)
		c.Next() // Pass on to the next-in-line middleware or handler
	}
}

func main() {
	// Load environment variables
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	// Get credentials from .env
	dbCredential := getCredentialsFromEnv()

	// Connect to database
	dbConn, err := Connect(dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate models
	modelsToManage := []interface{}{
		&models.User{},
		&models.Blacklist{},
		&models.Course{},
		&models.Module{},
		&models.Lesson{},
		&models.Test{},
		&models.QuestionEssay{},
		&models.QuestionOpt{},
		&models.AnswerOpt{},
	}

	// Drop existing tables (optional, for development/testing)
	for _, model := range modelsToManage {
		if err := dbConn.Migrator().DropTable(model); err != nil {
			log.Fatal("failed dropping table: " + err.Error())
		}
	}

	for _, model := range modelsToManage {
		if err := dbConn.AutoMigrate(model); err != nil {
			log.Fatal("failed creating table: " + err.Error())
		}
	}

	// Initialize services
	authService := services.NewAuthService(dbConn)

	// Initialize controllers
	userController := &controllers.UserController{DB: dbConn}
	authController := &controllers.AuthController{DB: dbConn}
	courseController := &controllers.CourseController{DB: dbConn}
	moduleController := &controllers.ModuleController{DB: dbConn}
	lessonController := &controllers.LessonController{DB: dbConn}
	testController := &controllers.TestController{DB: dbConn}
	questionOptController := &controllers.QuestionOptController{DB: dbConn}
	questionEssayController := &controllers.QuestionEssayController{DB: dbConn}
	answerOptController := &controllers.AnswerOptController{DB: dbConn}

	// Initialize Gin router
	router := gin.Default()

	// CORS setup using gin-contrib/cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Add the DB middleware
	router.Use(DBMiddleware(dbConn))

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Public authentication routes (no middleware)
	publicGroup := router.Group("/api")
	routes.AuthRoutes(publicGroup, authController)

	// Middleware
	router.Use(middlewares.AuthMiddleware(authService)) // Pass AuthService to AuthMiddleware

	// Set up routes
	apiGroup := router.Group("/api")
	routes.UserRoutes(apiGroup, userController, authService)
	routes.CourseRoutes(apiGroup, courseController, authService)
	routes.ModuleRoutes(apiGroup, moduleController, authService)
	routes.LessonRoutes(apiGroup, lessonController, authService)
	routes.TestRoutes(apiGroup, testController, authService)
	routes.QuestionOptRoutes(apiGroup, questionOptController, authService)
	routes.QuestionEssayRoutes(apiGroup, questionEssayController, authService)
	routes.AnswerOptRoutes(apiGroup, answerOptController, authService)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
