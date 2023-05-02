package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const TABLECITAS string = "citas"

func (con *Controllers) PerfilMedico(c *fiber.Ctx) error {

	return c.JSON("")
}

func (con *Controllers) CalendarioCitas(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var search search

	c.BodyParser(&search)

	db := con.Client

	coll, err := db.Collection(TABLECITAS)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	cur, err := coll.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []bson.D

	cur.All(ctx, &res)

	return c.JSON(Response{Message: res})
}

type CitaModificar struct {
	Fecha string `json:"fecha"`
	Hora  string `json:"hora"`
}

func (con *Controllers) GetCitas(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// var body map[string]interface{}

	// err := c.BodyParser(&body)

	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	// }

	db := con.Client

	coll, err := db.Collection(TABLECITAS)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// token := c.Cookies("token")

	// data := config.ExtractClaims(token)

	email := c.Params("email")

	//fmt.Println("email", email)
	// filter := bson.D{{
	// 	Key: "_id", Value: data["email"],
	// }}

	filterUser := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "users",      // tabla ala que esta relacionada
		"localField":   "idServicio", // local id o llave local
		"foreignField": "_id",        // llave la otra tabla
		"as":           "idServicio", // nombre que se llamara
	}}}

	filterServicio := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "users",  // tabla ala que esta relacionada
		"localField":   "idUser", // local id o llave local
		"foreignField": "_id",    // llave la otra tabla
		"as":           "idUser", // nombre que se llamara
	}}}

	filter := bson.D{{Key: "$match", Value: bson.M{
		"idServicio.email": email,
	}}}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filterUser, filterServicio, filter})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []models.ServicioCitas

	cur.All(ctx, &res)

	return c.JSON(Response{Message: res})
}

func (con *Controllers) ModificarCita(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id := c.Params("id")

	var cita CitaModificar
	c.BodyParser(&cita)

	db := con.Client

	coll, err := db.Collection(TABLECITAS)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	idEx, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": idEx}

	update := bson.M{"$set": bson.M{"fecha": cita.Fecha, "hora": cita.Hora}}

	_, err = coll.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}
