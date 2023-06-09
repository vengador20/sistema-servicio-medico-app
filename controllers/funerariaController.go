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

const (
	funeraria              string = "funeraria"
	TABLESERVICIOFUNERARIA string = "servicioFuneraria"
)

// func (con *Controllers) GetFuneraria(c *fiber.Ctx) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	var search search

// 	c.BodyParser(&search)

// 	db := con.Client

// 	coll, err := db.Collection(TABLEUSER)

// 	if err != nil {
// 		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
// 	}

// 	filterFuneraria := bson.D{{Key: "$lookup", Value: bson.M{
// 		"from":         "servicioFuneraria", // tabla ala que esta relacionada
// 		"localField":   "_id",               // local id o llave local
// 		"foreignField": "_id",               // llave la otra tabla
// 		"as":           "idGrupo",           // nombre que se llamara
// 	}}}

// 	// filterServicio := bson.D{{Key: "$lookup", Value: bson.M{
// 	// 	"from":         "serviciosFuneraria",
// 	// 	"localField":   "idGrupo",
// 	// 	"foreignField": "_id",
// 	// 	"as":           "idGrupo",
// 	// }}}

// 	cur, err := coll.Aggregate(ctx, mongo.Pipeline{filterFuneraria})

// 	if err != nil {
// 		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
// 	}

// 	var res models.UserService

// 	cur.All(ctx, &res)

// 	return c.JSON(Response{Message: res})
// }

func (con *Controllers) GetFunerariaByIdUpdate(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id := c.Params("id")

	db := con.Client

	coll, _ := db.Collection(TABLESERVICIOFUNERARIA)

	var res bson.M

	idEx, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	fmt.Println("id", id)

	filter := bson.D{{Key: "_id", Value: idEx}}

	err = coll.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetFunerariaById(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	id := c.Params("id")

	db := con.Client

	coll, _ := db.Collection(TABLEUSER)

	var res bson.M

	idEx, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.D{{Key: "_id", Value: idEx}}

	err = coll.FindOne(ctx, filter).Decode(&res)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetFuneraria(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	coll, _ := db.Collection(TABLESERVICIOFUNERARIA)

	funerariaSer := bson.D{{
		Key: "$lookup", Value: bson.M{
			"from":         "users",  // tabla ala que esta relacionada
			"localField":   "idUser", // local id o llave local
			"foreignField": "_id",    // llave la otra tabla
			"as":           "idUser", // nombre que se llamara
		},
	}}

	// filter := bson.D{{Key: "$and", Value: []bson.M{
	// 	{"tipo": 1},
	// 	{"servicio": funeraria},
	// },
	// }}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{funerariaSer})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []bson.M

	err = cur.All(ctx, &res)

	//fmt.Println(res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetFunerariaSearch(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	nombre := c.Params("nombre")

	fmt.Printf("nombre: %v\n", nombre)

	coll, _ := db.Collection(TABLESERVICIOFUNERARIA)

	funerariaSer := bson.D{{
		Key: "$lookup", Value: bson.M{
			"from":         "users",  // tabla ala que esta relacionada
			"localField":   "idUser", // local id o llave local
			"foreignField": "_id",    // llave la otra tabla
			"as":           "idUser", // nombre que se llamara
		},
	}}

	filterName := bson.D{{Key: "$match", Value: bson.M{
		"nombre": bson.M{"$regex": nombre},
	}}}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{funerariaSer, filterName})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []bson.M

	err = cur.All(ctx, &res)

	//fmt.Println(res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetFunerariaEmail(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	email := c.Params("email")

	coll, _ := db.Collection(TABLESERVICIOFUNERARIA)

	funerariaSer := bson.D{{
		Key: "$lookup", Value: bson.M{
			"from":         "users",  // tabla ala que esta relacionada
			"localField":   "idUser", // local id o llave local
			"foreignField": "_id",    // llave la otra tabla
			"as":           "idUser", // nombre que se llamara
		},
	}}

	filterUS := bson.D{{Key: "$match", Value: bson.M{
		"idUser.email": email,
	}}}

	// filter := bson.D{{Key: "$and", Value: []bson.M{
	// 	{"tipo": 1},
	// 	{"servicio": funeraria},
	// },
	// }}

	cur, err := coll.Aggregate(ctx, mongo.Pipeline{funerariaSer, filterUS})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []bson.M

	err = cur.All(ctx, &res)

	//fmt.Println(res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) CrearServicioFuneraria(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var servicio models.ServiciosFuneraria

	err := c.BodyParser(&servicio)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error"})
	}

	errors, err := config.Validate(servicio)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	db := con.Client

	cur, err := db.Collection(TABLESERVICIOFUNERARIA)

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	insert := bson.M{
		"idUser": servicio.IdUser,
		"nombre": servicio.Nombre,
		"costo":  servicio.Costo,
	}

	_, err = cur.InsertOne(ctx, insert)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}

func (con *Controllers) ModificarServicioFuneraria(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var servicio models.ServiciosFuneraria

	err := c.BodyParser(&servicio)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	servicio.IdUser = primitive.NewObjectID()

	errors, err := config.Validate(servicio)

	//los mensajes de error cuando los campos son requeridos etc...
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(Response{Message: "error", Errors: errors})
	}

	var serId map[string]interface{}

	c.BodyParser(&serId)

	//fmt.Println(serId, servicio, serId["id"])

	db := con.Client

	cur, err := db.Collection(TABLESERVICIOFUNERARIA)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	id, err := primitive.ObjectIDFromHex(serId["id"].(string))
	//fmt.Println(id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	//fmt.Println(servicio.Nombre)
	update := bson.D{{Key: "$set", Value: bson.M{"nombre": servicio.Nombre, "costo": servicio.Costo}}}

	filter := bson.M{"_id": id}

	_, err = cur.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}

func (con *Controllers) EliminarServicioFuneraria(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	//var serId map[string]interface{}

	//c.BodyParser(&serId)

	idDelete := c.Params("id")

	db := con.Client

	cur, err := db.Collection(TABLESERVICIOFUNERARIA)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	id, err := primitive.ObjectIDFromHex(idDelete)
	//fmt.Println(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	_, err = cur.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: "exito"})
}
