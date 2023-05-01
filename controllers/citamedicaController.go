package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/config"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (con *Controllers) AgendarCitaMedica(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	var body models.Cita

	err := c.BodyParser(&body)
	fmt.Println("erorr")

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	fmt.Println("error")

	errors, err := config.Validate(body)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	fmt.Println("error")

	coll, err := db.Collection(TABLECITAS)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	fmt.Println("guardar")

	IdUser, _ := primitive.ObjectIDFromHex(body.IdUser)
	IdServicio, _ := primitive.ObjectIDFromHex(body.IdServicio)

	insert := bson.M{
		"fecha":          body.Fecha,
		"hora":           body.Hora,
		"pacienteNombre": body.NombreCompleto,
		"telefono":       body.Telefono,
		"alergias":       body.Alergias,
		"idUser":         IdUser,     //id del servicio usuario
		"idServicio":     IdServicio, //id del servicio medico
	}

	coll.InsertOne(ctx, insert)

	return c.JSON(Response{Message: "La cita se ha agendado exitosamente"})
}
