package authusercontroller

import (
	"fmt"
	"time"

	"github.com/Dandy-CP/go-fiber-portfolio/config"
	"github.com/Dandy-CP/go-fiber-portfolio/middleware"
	"github.com/Dandy-CP/go-fiber-portfolio/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthLogin(c *fiber.Ctx) error {
	var userInDB models.User
	var body models.User

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if config.DB.Where("Username = ?", body.Username).Find(&userInDB).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": fiber.StatusBadRequest,
			"message": "Wrong Password Or Username",
		})
	}


	checkStatusHash := 	middleware.CheckPasswordHash(body.Password, userInDB.Password)
	if !checkStatusHash {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": fiber.StatusBadRequest,
			"message": "Wrong Password Or Username",
		})
	}

	setup, _ := config.LoadConfig(".")

	tokenByte := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	claims["sub"] = userInDB.ID
	claims["exp"] = now.Add(setup.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	tokenString, err := tokenByte.SignedString([]byte(setup.JwtSecret))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"status": "fail",
			"message": fmt.Sprintf("generating JWT Token failed: %v", err),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		MaxAge:   setup.JwtMaxAge * 60,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": userInDB.ID,
		"username": userInDB.Username,
		"access_token": tokenString,
	})
}

func AuthSignUp(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	if config.DB.Where("Username = ?", user.Username).Find(&user).RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status": fiber.StatusBadRequest,
			"message": "Username Has been used",
		})
	}

	valueHash, errHasing := middleware.HashPassword(user.Password)
	if errHasing != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  fiber.StatusBadRequest,
			"Message": "error",
			"Data":    errHasing.Error(),
		})
	}

	newUser := models.User{
		Username: user.Username,
		Password: valueHash,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  fiber.StatusInternalServerError,
			"Message": "error",
			"Data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status": fiber.StatusOK,
		"Message": "Success Register",
	})
}

func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success log out",
	})
}