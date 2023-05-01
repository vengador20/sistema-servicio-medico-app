package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
)

func ValidateJwt(c *fiber.Ctx) error {
	cookie := c.Cookies("token")

	fmt.Printf("cookie: %v\n", cookie)

	fmt.Println(cookie)

	// data := config.ExtractClaims(cookie)

	// fmt.Println("token")
	// fmt.Println(data)

	status, err := config.VerifyToken(string(cookie))

	if err != nil {
		return c.SendStatus(400)
	}

	if !status {
		return c.SendStatus(400)
	}

	return c.Next()
}
