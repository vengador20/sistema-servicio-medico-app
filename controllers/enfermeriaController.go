package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const enfermeros = "enfermeros"

func (con *Controllers) GetEnfermerosSearch(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	nombre := c.Params("nombre")

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.D{{Key: "$and",
		Value: []bson.M{
			{"tipo": 1},
			{"servicio": enfermeros},
			{"nombreCompleto": bson.M{"$regex": nombre}},
		},
	}}

	cur, err := coll.Find(ctx, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []models.UserService

	err = cur.All(ctx, &res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetEnfermeros(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	filter := bson.D{{Key: "$and",
		Value: []bson.M{
			{"tipo": 1},
			{"servicio": enfermeros},
		},
	}}

	cur, err := coll.Find(ctx, filter)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res []models.UserService

	err = cur.All(ctx, &res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})
}

func (con *Controllers) GetEnfermerosById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	db := con.Client

	data := c.Params("id")

	//var body search

	//c.BodyParser(&body)

	coll, err := db.Collection(TABLEUSER)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	var res models.UserService

	id, err := primitive.ObjectIDFromHex(data)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	err = coll.FindOne(ctx, bson.M{"_id": id}).Decode(&res)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(Response{Message: "error"})
	}

	return c.JSON(Response{Message: res})

}
