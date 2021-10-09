package handlers

import (
	"money_tracker_go/database"
	"money_tracker_go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense
	database.Db.Find(&expenses)
	return c.JSON(expenses)
}

func InsertExpense(c *fiber.Ctx) error {
	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var err error
	expense.Id, err = uuid.NewV4()
	if err != nil {
		panic(err.Error())
	}
	database.Db.Create(&expense)

	return c.JSON(expense)
}
