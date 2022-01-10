package main

import (
	"fmt"
	"os"

	"github.com/ajuljulian/react-jwt-go/handler"
	"github.com/ajuljulian/react-jwt-go/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Set up the database
	db := initialMigration()

	// Initialize handler
	h := &handler.Handler{DB: db}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for signup and login requests
			if c.Path() == "/api/auth/register" || c.Path() == "/api/auth/login" || c.Path() == "/api/test/all" {
				return true
			}
			return false
		},
	}))

	// Routes
	e.POST("/api/auth/register", h.Register)
	e.POST("/api/auth/login", h.Login)

	e.GET("/api/test/all", h.PublicBoard)
	e.GET("/api/test/user", h.UserBoard)
	e.GET("/api/test/moderator", h.ModeratorBoard)
	e.GET("/api/test/admin", h.AdminBoard)

	e.Logger.Fatal(e.Start(":1323"))
}

func initialMigration() *gorm.DB {
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	database := os.Getenv("PGDATABASE")
	password := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, database, port)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	return db
}
