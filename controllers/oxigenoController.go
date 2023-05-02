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
	"go.mongodb.org/mongo-driver/mongo"
)

const oxigeno = "oxigeno"

// func (con *Controllers) GetOxigeno(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	db := con.Client

// 	coll, err := db.Collection(TABLEUSER)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
// 	}

// 	filter := bson.D{{Key: "$and",
// 		Value: []bson.M{
// 			{"tipo": 1},
// 			{"servicio": oxigeno},
// 		},
// 	}}

// 	cur, err := coll.Find(ctx, filter)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
// 	}

// 	var res []models.UserService

// 	err = cur.All(ctx, &res)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
// 	}

// 	return c.JSON(Response{Message: res})
// }

func (con *Controllers) GetOxigeno(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	coll, err := db.Collection(oxigeno)

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

	var res []models.OxigenoModel

	err = cur.All(ctx, &res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetOxigenoEmail(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	email := c.Params("email")

	coll, err := db.Collection(oxigeno)

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

	filterUS := bson.D{{Key: "$match", Value: bson.M{
		"idServicio.email": email,
	}}}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filter, filterUS})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []models.OxigenoModel

	err = cur.All(ctx, &res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetOxigenoByid(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	id := c.Params("id")

	coll, err := db.Collection(oxigeno)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	idServicio, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}
	//fmt.Printf("idServicio: %v\n", idServicio)
	filter := bson.D{{
		Key: "_id", Value: idServicio,
	}}

	var res models.ServiciOxigeno

	err = coll.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	fmt.Println(res)

	return c.JSON(Response{Message: res})
}

func (con *Controllers) PerfilOxigeno(c *fiber.Ctx) error {
	//fmt.Println("error")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// var body map[string]interface{}

	// err := c.BodyParser(&body)

	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	// }

	token := c.Cookies("token")

	data := config.ExtractClaims(token)

	db := con.Client

	coll, err := db.Collection(oxigeno)

	//fmt.Println("oxigeno")

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	// filter := bson.D{{
	// 	Key: "_id", Value: idServicio,
	// }}

	filterServicio := bson.D{{Key: "$lookup", Value: bson.M{
		"from":         "users",      // tabla ala que esta relacionada
		"localField":   "idServicio", // local id o llave local
		"foreignField": "_id",        // llave la otra tabla
		"as":           "idServicio", // nombre que se llamara
	}}}

	filter := bson.D{{Key: "$match", Value: bson.M{
		"idServicio.email": data["email"],
	}}}

	//fmt.Println("agregate")

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filterServicio, filter})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []bson.M

	cur.All(ctx, &res)

	return c.JSON(Response{Message: res})
}

func (con *Controllers) CrearOxigeno(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	var body models.ServiciOxigeno

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	cur, err := db.Collection(oxigeno)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	insert := bson.M{
		"nombre":     body.Nombre,
		"costo":      body.Costo,
		"tipo":       body.Tipo,
		"idServicio": body.IdServicio,
	}

	_, err = cur.InsertOne(ctx, insert)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}

func (con *Controllers) ModificarOxigeno(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var oxigenoModels models.ServiciOxigeno

	err := c.BodyParser(&oxigenoModels)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	id := c.Params("id")

	idOxigeno, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	db := con.Client

	coll, err := db.Collection(oxigeno)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.M{"_id": idOxigeno}

	update := bson.D{{Key: "$set", Value: bson.M{"nombre": oxigenoModels.Nombre, "costo": oxigenoModels.Costo,
		"tipo": oxigenoModels.Tipo}}}

	_, err = coll.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}

func (con *Controllers) EliminarOxigeno(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id := c.Params("id")

	idOxigeno, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	db := con.Client

	coll, err := db.Collection(oxigeno)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.M{"_id": idOxigeno}

	_, err = coll.DeleteOne(ctx, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}
