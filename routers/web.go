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

func IsFuneraria(url string) map[string]any {
	var js map[string]any

	resServicio, err := http.Get(url)

	if err != nil {
		js["message"] = ""
		return js
	}

	bodySer, err := io.ReadAll(resServicio.Body)
	fmt.Println("nil")

	if err != nil {
		js["message"] = ""
		return js
	}

	err = json.Unmarshal(bodySer, &js)

	if err != nil {
		js["message"] = ""
		return js
	}

	fmt.Printf("js: %v\n", js)
	return js
}

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
		// wg.Add(1)

		// go func() {
		// 	defer wg.Done()

		// 	res, err := http.Get("http://localhost:3000/api/servicio-funeraria")

		// 	if err != nil {
		// 		js["message"] = ""
		// 		return
		// 	}

		// 	body, err := io.ReadAll(res.Body)

		// 	if err != nil {
		// 		js["message"] = ""
		// 		return
		// 	}
		// 	fmt.Println("hola")
		// 	err = json.Unmarshal(body, &js)

		// 	if err != nil {
		// 		js["message"] = ""
		// 		return
		// 	}

		// 	fmt.Printf("js: %v\n", js)

		// 	return
		// 	// return c.Render("oxigeno", fiber.Map{
		// 	// 	"message": js["message"],
		// 	// })
		// }()

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

		//wg.Wait()
		var js map[string]any

		if user.Message.Servicio == "funeraria" {
			js = IsFuneraria("http://localhost:3000/api/servicio-funeraria/" + user.Message.Email)
		} else if user.Message.Servicio == "oxigeno" {
			js = IsFuneraria("http://localhost:3000/api/oxigeno/" + user.Message.Email)
		} else if user.Message.Servicio == "medico" {
			js = IsFuneraria("http://localhost:3000/api/perfil/servicio-medico/" + user.Message.Email)
		} else if user.Message.Servicio == "enfermeros" {
			js = IsFuneraria("http://localhost:3000/api/perfil/servicio-enfermeria/" + user.Message.Email)
		}

		fmt.Println(js)

		return c.Render("perfilprestador", fiber.Map{
			"telefono":       user.Message.Telefono,
			"email":          user.Message.Email,
			"servicio":       user.Message.Servicio,
			"nombreCompleto": user.Message.NombreCompleto,
			"nombreNegocio":  user.Message.NombreNegocio,
			"message":        js["message"],
		})
	})

	router.Get("/perfil/funeraria/update/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		res, err := http.Get("http://localhost:3000/api/servicio-funeraria/update-one/" + id)

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
		fmt.Println(js)

		return c.Render("servicio/update/funeraria", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/perfil/funeraria/create", func(c *fiber.Ctx) error {

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
		fmt.Println(user.Message)
		if user.Message.Tipo == 0 {
			return c.Render("404", "")
		}

		return c.Render("servicio/form/funeraria", fiber.Map{
			"message": user.Message.Id.Hex(),
		})
	})

	router.Get("/perfil/oxigeno/create", func(c *fiber.Ctx) error {

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
		fmt.Println(user.Message)
		if user.Message.Tipo == 0 {
			return c.Render("404", "")
		}

		return c.Render("servicio/form/oxigeno", fiber.Map{
			"message": user.Message.Id.Hex(),
		})
	})

	router.Get("/perfil/oxigeno/update/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		res, err := http.Get("http://localhost:3000/api/oxigeno/" + id)

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
			c.Render("oxigeno", "")
		}
		fmt.Println(js)

		return c.Render("servicio/update/oxigeno", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/perfil/medico/update/:id", func(c *fiber.Ctx) error {

		id := c.Params("id")

		res, err := http.Get("http://localhost:3000/api/cita-medica/" + id)

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
			c.Render("medico", "")
		}
		fmt.Println(js)

		return c.Render("servicio/update/medico", fiber.Map{
			"message": js["message"],
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

		res, err := http.Get("http://localhost:3000/api/enfermeros")

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
			c.Render("funeraria", fiber.Map{
				"message": "",
			})
		}

		return c.Render("enfermeria", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/oxigeno", func(c *fiber.Ctx) error {

		res, err := http.Get("http://localhost:3000/api/oxigeno")

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
			c.Render("oxigeno", fiber.Map{
				"message": "",
			})
		}

		return c.Render("oxigeno", fiber.Map{
			"message": js["message"],
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

	router.Get("/search/oxigeno", func(c *fiber.Ctx) error {
		nombre := c.FormValue("nombre")

		res, err := http.Get("http://localhost:3000/api/oxigeno/search/" + nombre)

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
			c.Render("oxigeno", fiber.Map{
				"message": "",
			})
		}

		return c.Render("servicio/search/oxigeno", fiber.Map{
			"message": js["message"],
		})

		//return c.Render("servicio/search/oxigeno", "")
	})

	router.Get("/search/medico", func(c *fiber.Ctx) error {
		nombre := c.FormValue("nombre")

		res, err := http.Get("http://localhost:3000/api/medico/search/" + nombre)

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
			c.Render("medico", fiber.Map{
				"message": "",
			})
		}

		return c.Render("servicio/form/medico", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/search/funeraria", func(c *fiber.Ctx) error {
		nombre := c.FormValue("nombre")

		res, err := http.Get("http://localhost:3000/api/servicio-funeraria/search/" + nombre)

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
			c.Render("funeraria", fiber.Map{
				"message": "",
			})
		}

		return c.Render("servicio/search/funeraria", fiber.Map{
			"message": js["message"],
		})
	})

	router.Get("/search/enfermeria", func(c *fiber.Ctx) error {

		nombre := c.FormValue("nombre")

		res, err := http.Get("http://localhost:3000/api/enfermeros/search/" + nombre)

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
			c.Render("funeraria", fiber.Map{
				"message": "",
			})
		}

		return c.Render("servicio/search/enfermeria", fiber.Map{
			"message": js["message"],
		})
	})
}
