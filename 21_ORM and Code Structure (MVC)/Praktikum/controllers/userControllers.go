package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"praktikum/models"

	"github.com/labstack/echo/v4"
)

// Get all user controller
func GetUserControllers(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all user",
		"user":   users,
	})
}

// Get user by id controller
func GetUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user, err := database.GetUserById(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get user by id",
		"user":   user,
	})
}

// Create user controller
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	user, err := database.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success create user",
		"user":   user,
	})
}

// Update user by id controller
func UpdateUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	user := models.User{}
	c.Bind(&user)

	user, err := database.UpdateUser(user, UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update user by id",
		"user":   user,
	})
}

// Delete user by id controller-
func DeleteUserByIdController(c echo.Context) error {
	UserId := c.Param("id")

	_, err := database.DeleteUser(UserId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user by id",
	})
}
