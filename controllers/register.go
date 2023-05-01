package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func registerUser(c *fiber.Ctx, ctx context.Context, con *Controllers) error {

	var user models.UserRegister

	err := c.BodyParser(&user)
	fmt.Printf("user: %v\n", user)

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

	//verificar que el usuario no existe
	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var userRes bson.M

	filter := bson.M{"email": user.Email}
	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//si es nulo  usuario ya existe
	if err == nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Usuario ya existe"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// asignamos el password anterior
	// al password encriptado
	user.Password = string(password)

	_, err = coll.InsertOne(ctx, user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "Usuario creado"})
}

func registerService(c *fiber.Ctx, ctx context.Context, con *Controllers) error {

	var user models.ServicioRegister

	err := c.BodyParser(&user)
	// fmt.Printf("user: %v\n", user)
	// fmt.Println(user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: err.Error()})
	}

	//fmt.Println(user.Telefono)

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

	//verificar que el usuario no existe
	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var userRes bson.M

	filter := bson.M{"email": user.Email}
	err = coll.FindOne(ctx, filter).Decode(&userRes)

	//si es nulo  usuario ya existe
	if err == nil {
		return c.Status(http.StatusUnauthorized).JSON(Response{Message: "Usuario ya existe"})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// asignamos el password anterior
	// al password encriptado
	user.Password = string(password)

	//user.
	inser := bson.M{
		"nombreCompleto": user.NombreCompleto,
		"nombreNegocio":  user.NombreNegocio,
		"telefono":       user.Telefono,
		"email":          user.Email,
		"tipo":           user.Tipo,
		"servicio":       user.Servicio,
		"latitud":        user.Latitud,
		"longitud":       user.Longitud,
		"cedula":         user.Cedula,
		"password":       user.Password,
	}

	_, err = coll.InsertOne(ctx, inser)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "Usuario creado"})
}
