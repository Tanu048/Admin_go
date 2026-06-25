package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/Tanu048/Admin_go/database"
	"github.com/Tanu048/Admin_go/models"
)

var jwtSecret = []byte("super-secret-isparc-key") // Change later
type AdminLoginInput struct {
	AdminID  string `json:"admin_id"`
	Password string `json:"password"`
}

func LoginAdmin(c *fiber.Ctx) error {
	var input AdminLoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	var mentor models.Admin

	// Search admin by Admin ID
	if err := database.DB.Where("admin_id = ?", input.AdminID).First(&mentor).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid Admin ID or Password"})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(mentor.PasswordHash), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid Admin ID or Password"})
	}

	// Create a JWT token
	claims := jwt.MapClaims{
		"id":             mentor.ID,
		"admin_id":       mentor.AdminID,
		"is_super_admin": mentor.IsSuperAdmin,
		"exp":            time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not generate login token"})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"token":   t,
	})
}
