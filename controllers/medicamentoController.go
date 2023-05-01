package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	TABLEFARMACIA string = "farmacia"
	medicamento   string = "medicamento"
)

type search struct {
	Data string `json:"data"`
}

func (con *Controllers) ConsultaMedicamento(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	nombre := c.Params("name")

	//buscar en la base de datos
	db := con.Client

	coll, err := db.Collection(medicamento)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filterServicio := bson.D{{
		Key: "$lookup", Value: bson.M{
			"from":         "users",      // tabla ala que esta relacionada
			"localField":   "idServicio", // local id o llave local
			"foreignField": "_id",        // llave la otra tabla
			"as":           "idServicio", // nombre que se llamara
		},
	}}

	//la i significa que no es sensible
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "nombre", Value: bson.M{"$regex": nombre, "$options": "i"}}}}}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filterServicio, filter})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []models.MedicamentoModel

	cur.All(ctx, &res)

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetMedicamentos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	//buscar en la base de datos
	db := con.Client

	coll, err := db.Collection(medicamento)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.D{{
		Key: "$lookup", Value: bson.M{
			"from":         "users",      // tabla ala que esta relacionada
			"localField":   "idServicio", // local id o llave local
			"foreignField": "_id",        // llave la otra tabla
			"as":           "idServicio", // nombre que se llamara
		},
	}}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filter})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var medicamentos []models.MedicamentoModel

	cur.All(ctx, &medicamentos)

	return c.JSON(Response{Message: medicamentos})
}

func (con *Controllers) CrearMedicamento(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	var body models.Medicamento

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//buscar en la base de datos
	db := con.Client

	coll, err := db.Collection(medicamento)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	document := bson.M{
		"nombre":     body.Nombre,
		"imagen":     body.Imagen,
		"precio":     body.Precio,
		"domicilio":  body.Domicilio,
		"idServicio": body.IdServicio,
	}

	_, err = coll.InsertOne(ctx, document)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}
