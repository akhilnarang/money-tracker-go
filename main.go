package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"money_tracker_go/database"
	"money_tracker_go/handlers"
	"os"
)

func InitDb() {
	var err error

	dsn := os.Getenv("DATABASE_URI")

	if dsn == "" {
		panic("Please provide a proper database URI")
	}

	database.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	app := fiber.New(fiber.Config{AppName: "Money Tracker", Prefork: true})

	app.Use(recover.New())

	InitDb()

	app.Get("/", handlers.Root)

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/", handlers.GetExpenses)

	apiV1.Get("/:id", handlers.GetExpenseById)

	apiV1.Post("/", handlers.InsertExpense)

	apiV1.Delete("/:id", handlers.DeleteExpense)

	log.Fatal(app.Listen(":8000"))

}
