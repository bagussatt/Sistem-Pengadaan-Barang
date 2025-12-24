package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"purchasing-backend/config"
	"purchasing-backend/models"
	"purchasing-backend/utils"
)

type RegisterRequest struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password123"`
	Role     string `json:"role" example:"admin"`
}

type LoginRequest struct {
	Username string `json:"username" example:"johndoe"`
	Password string `json:"password" example:"password123"`
}

// Register godoc
// @Summary Register new user
// @Description Register a new user with username, password and role
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Register Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username" example:"johndoe"`
		Password string `json:"password" example:"password123"`
		Role     string `json:"role" example:"admin"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Username == "" || req.Password == "" || req.Role == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Username, password, and role are required"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Username already exists"})
	}

	return c.JSON(fiber.Map{
		"message": "Register success",
	})
}

// Login godoc
// @Summary Login user
// @Description Login with username and password to get JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login Request"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username" example:"johndoe"`
		Password string `json:"password" example:"password123"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if err := config.DB.
		Where("username = ?", req.Username).
		First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Invalid username or password"})
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": "Invalid username or password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
