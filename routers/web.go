package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
)

type response struct {
	Message models.ServicioRegister `json:"message"`
	Errors  []string                `json:"errors"`
}

// type responseJson struct {
// 	Message map[string]interface{} `json:"message"`
// 	Errors  []string               `json:"errors"`
// }

var wg sync.WaitGroup

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

	router.Get("/medico", func(c *fiber.Ctx) error {

		res, err := http.Get("http://localhost:3000/api/medico")

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		//var js responseJson

		var rs map[string]any

		err = json.Unmarshal(body, &rs)

		//err = json.Unmarshal(body, &js)

		if err != nil {
			c.Render("medico", "")
		}

		//fmt.Println(js)

		return c.Render("medico", fiber.Map{
			"message": rs["message"],
		})
	})

	router.Get("/medico/cita/:id", func(c *fiber.Ctx) error {

		type response struct {
			Message models.UserService `json:"message"`
			Errors  []string           `json:"errors"`
		}

		var idUser string

		wg.Add(1)
		go func() {
			cookie := c.Cookies("token")

			token := config.ExtractClaims(cookie)

			defer wg.Done()

			if token["email"] == nil || token["email"] == "" {
				idUser = ""
				return
			}

			res, err := http.Get("http://localhost:3000/api/user/" + token["email"].(string))

			if err != nil {
				idUser = ""
				return
			}

			body, err := io.ReadAll(res.Body)

			if err != nil {
				idUser = ""
				return
			}

			var user response

			err = json.Unmarshal(body, &user)

			if err != nil {
				idUser = ""
				return
			}

			idUser = user.Message.Id.Hex() //String()

		}()

		id := c.Params("id")

		res, err := http.Get("http://localhost:3000/api/medico/" + id)

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		var js response

		err = json.Unmarshal(body, &js)

		if err != nil {
			c.Render("404", "")
		}

		// fmt.Println(js.Message["telefono"].(int))
		// fmt.Println("telefono")

		// fmt.Println(js)
		// fmt.Println(reflect.TypeOf(js.Message.Telefono))

		// json.Unmarshal([]byte(js.Message))

		s := structs.Map(&js.Message)

		s["_id"] = js.Message.Id.Hex() //String()

		fmt.Println(s["_id"])
		fmt.Println(js.Message.Id.Hex())

		fmt.Println(s["Servicio"])

		wg.Wait()

		return c.Render("datosericioMF", fiber.Map{
			"message": s,
			"idUser":  idUser,
			//"int":      func(str any) int { return str.(int) },
			//"telefono": js.Message["telefono"].(int),
		})
	})

	router.Get("/funeraria", func(c *fiber.Ctx) error {

		res, err := http.Get("http://localhost:3000/api/servicio-funeraria")

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		var js map[string]any

		err = json.Unmarshal(body, &js)

		if err != nil {
			c.Render("funeraria", "")
		}

		return c.Render("funeraria", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/funeraria/servicio/:id", func(c *fiber.Ctx) error {

		type response struct {
			Message models.UserService `json:"message"`
			Errors  []string           `json:"errors"`
		}

		id := c.Params("id")

		res, err := http.Get("http://localhost:3000/api/servicio-funeraria/" + id)

		if err != nil {
			fmt.Println(err.Error())
		}

		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println("error")
		}

		var js response

		err = json.Unmarshal(body, &js)

		if err != nil {
			c.Render("404", "")
		}

		// fmt.Println(js.Message["telefono"].(int))
		// fmt.Println("telefono")

		// fmt.Println(js)
		// fmt.Println(reflect.TypeOf(js.Message.Telefono))

		// json.Unmarshal([]byte(js.Message))

		s := structs.Map(&js.Message)

		s["_id"] = js.Message.Id.Hex() //String()

		// fmt.Println(s["_id"])
		// fmt.Println(js.Message.Id.Hex())

		// fmt.Println(s)

		return c.Render("datoservicio", fiber.Map{
			"message": s, //js["message"],
		})
	})

	router.Get("/enfermeria", func(c *fiber.Ctx) error {
		return c.Render("enfermeria", "")
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
