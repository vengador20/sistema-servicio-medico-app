package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Message interface {
	string | interface{}
}

type Response struct {
	Message Message  `json:"message"`
	Errors  []string `json:"errors"`
}

const TABLEUSER string = "users"

type register interface {
	models.UserRegister | models.ServicioRegister
}

type MyStruct struct {
	Name string
	Age  int
}

// terminado
func (con *Controllers) Login(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.UserLogin

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//message de errores
	errors, err := config.ValidateUser(&user)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	// db := database.Mongodb{
	// 	Client: con.Client,
	// }
	db := con.Client

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error"})
	}

	var userRes models.UserLogin

	//filtrar para la busqueda del usuario si existe
	filter := bson.D{{Key: "email", Value: user.Email}}

	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//usuario no existe
	if err != nil {
		errors = []string{"Correo o contraseña inválido"}
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "error", Errors: errors})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(user.Password))

	//contraseña invalida
	if err != nil {
		errors = []string{"Correo o contraseña inválido"}
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "error", Errors: errors})
	}

	//craer token
	token, err := config.NewToken(user.Email, userRes.Tipo)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error"})
	}

	//crear token con la cookie
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), //expiracion 24 horas
		HTTPOnly: true,
		Secure:   true,
		Path:     "/",
		//cookie.Expires = time.Now().Add(24 * time.Hour)
		SameSite: "none",
	}

	//insertamos la cookie al usuario
	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(Response{Message: "exito"})
}

func (con *Controllers) RegisterUser(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var mp map[string]interface{}

	c.BodyParser(&mp)

	//fmt.Println(mp["tipo"])

	var tipo float64 = mp["tipo"].(float64)
	var err error

	// 0 es un paciente y 1 es servicio	switch tipo {
	switch tipo {
	case 0:
		err = registerUser(c, ctx, con)
	case 1:
		err = registerService(c, ctx, con)
	}

	return err
}

// restablecer password
func (con *Controllers) ResetPassword(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var user models.UserNewPassword

	err := c.BodyParser(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// db := database.Mongodb{
	// 	Client: con.Client,
	// }
	db := con.Client

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var userRes bson.M

	filter := bson.M{"email": user.Email}
	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//usuario no existe
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Usuario no existe"})
	}

	filterPassword := bson.M{"email": user.Email}

	update := bson.M{"password": user.Password}

	_, err = coll.UpdateOne(ctx, filterPassword, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//enviar correo
	//smtp.PlainAuth("","vengadorba6@gmail.com","tobimoto2000","")

	return c.Status(http.StatusOK).JSON(Response{Message: "Contraseña actualizada"})
}

func (con *Controllers) GetUserEmail(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	email := c.Params("email")

	db := con.Client

	coll, err := db.Collection("users")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.D{{Key: "email", Value: email}}

	var res bson.M

	err = coll.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "usuario no existe"})
	}

	delete(res, "password")

	return c.JSON(Response{Message: res})
}
