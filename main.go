package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"github.com/vengador20/sistema-servicios-medicos/controllers"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"github.com/vengador20/sistema-servicios-medicos/routers"
)

func main() {

	engine := html.New("./view", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./public")

	app.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir("./public/img/servicios"),
		Browse: true,
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     "*",
	// }))

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:5173",
	}))

	app.Use(recover.New())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	db, err := database.Conn()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Disconnect(ctx)

	var dbInteface database.DatabaseMongodb = db

	rt := routers.Router{
		Db: dbInteface,
	}

	app.Post("/file/save", func(c *fiber.Ctx) error {

		var body map[string]interface{}
		err := c.BodyParser(&body)

		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		err = controllers.SaveImagen(body["nombre"].(string), body["file"].(string))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "error",
			})
		}

		return c.JSON(fiber.Map{
			"message": "exito",
		})
	})
	//utilizar middleware personalizado
	//valida si el jwt no es modificado
	//app.Use("/api", middleware.ValidateJwt)

	app.Route("/api", rt.Router)

	app.Route("/", routers.Web)

	app.Listen(":3000")
}
