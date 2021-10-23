package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"money_tracker_go/database"
	"money_tracker_go/handlers"
	"os"
)

var port = flag.Int("port", 8000, "Port to listen on")

func InitDb() {
	var err error

	// Fetch our DB URI from the environment
	dsn := os.Getenv("DATABASE_URI")

	// Ensure its set
	if dsn == "" {
		panic("Please provide a proper database URI")
	}

	// Open a connection to the database
	database.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// Parse flags
	flag.Parse()

	// Create our fiber app
	app := fiber.New(fiber.Config{AppName: "Money Tracker", Prefork: true})

	// Use recover middleware
	app.Use(recover.New())

	// Initialize DB connection
	InitDb()

	// Start declaring our routes
	app.Get("/", handlers.Root)

	// Create a group for /api/v1 prefix
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/", handlers.GetExpenses)

	apiV1.Get("/:id", handlers.GetExpenseById)

	apiV1.Post("/", handlers.InsertExpense)

	apiV1.Delete("/:id", handlers.DeleteExpense)

	// Start our fiber app
	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
