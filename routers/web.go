package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
)

type response struct {
	Message models.ServicioRegister `json:"message"`
	Errors  []string                `json:"errors"`
}

func Web(router fiber.Router) {
	router.Get("/", func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")

		//fmt.Println(c.Cookies("token"))

		token := config.ExtractClaims(cookie)

		if token["email"] == "" || token["email"] == nil {
			return c.Render("index", fiber.Map{
				"Login": false,
			})
		}
		res, err := http.Get("http://localhost:3000/api/user/" + token["email"].(string))

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		var user response

		err = json.Unmarshal(body, &user)

		if err != nil {
			fmt.Println("error")
		}
		return c.Render("index", fiber.Map{
			"Login":   true,
			"Message": user.Message,
			"string":  func(str any) int { return str.(int) },
		})
	})

	router.Get("/user/profile", func(c *fiber.Ctx) error {

		cookie := c.Cookies("token")

		token := config.ExtractClaims(cookie)

		if token["email"] == nil || token["email"] == "" {
			return c.Render("404", "")
		}

		res, err := http.Get("http://localhost:3000/api/user/" + token["email"].(string))

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		var user response

		err = json.Unmarshal(body, &user)

		if err != nil {
			fmt.Println("error")
		}

		if user.Message.Tipo == 0 {
			return c.Render("404", "")
		}

		return c.Render("perfilprestador", fiber.Map{
			"telefono":       user.Message.Telefono,
			"email":          user.Message.Email,
			"servicio":       user.Message.Servicio,
			"nombreCompleto": user.Message.NombreCompleto,
			"nombreNegocio":  user.Message.NombreNegocio,
		})
	})

	router.Get("/login/servicio", func(c *fiber.Ctx) error {
		return c.Render("login1", "")
	})

	router.Get("/login/cliente", func(c *fiber.Ctx) error {
		return c.Render("login2", "")
	})

	router.Get("/register/servicio", func(c *fiber.Ctx) error {
		return c.Render("formulario", "")
	})

	router.Get("/register/cliente", func(c *fiber.Ctx) error {
		return c.Render("formcliente", "")
	})
}
