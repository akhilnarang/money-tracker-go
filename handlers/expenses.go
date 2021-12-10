package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"money_tracker_go/database"
	"money_tracker_go/models"
)

func GetExpenses(c *fiber.Ctx) error {
	var expenses []models.Expense
	database.Db.Find(&expenses)
	return c.JSON(expenses)
}

func GetExpenseById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var expense models.Expense
	result := database.Db.First(&expense, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("Expense with ID %s not found", id)})
	}
	return c.JSON(expense)
}

func InsertExpense(c *fiber.Ctx) error {
	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	expense.Id = uuid.New()
	database.Db.Create(&expense)

	return c.Status(fiber.StatusCreated).JSON(expense)
}

func DeleteExpense(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	result := database.Db.Delete(&models.Expense{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("Expense with ID %s not found", id)})
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Deleted expense with ID %s", id)})
}

func UpdateExpense(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	var expense models.Expense
	result := database.Db.First(&expense, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": fmt.Sprintf("Expense with ID %s not found", id)})
	}

	expenseUpdate := new(models.ExpenseUpdate)
	if err := c.BodyParser(expenseUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if expenseUpdate.PaymentMethod != "" {
		expense.PaymentMethod = expenseUpdate.PaymentMethod
	}
	if expenseUpdate.Amount != 0 {
		expense.Amount = expenseUpdate.Amount
	}
	if expenseUpdate.Description != "" {
		expense.Description = expenseUpdate.Description
	}
	database.Db.Save(&expense)
	return c.JSON(expense)
}
