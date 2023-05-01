package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
	"go.mongodb.org/mongo-driver/bson"
)

const centroSalud = "centroSalud"

func (con *Controllers) GetCentroSalud(c *fiber.Ctx) error {
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
			{"servicio": centroSalud},
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
