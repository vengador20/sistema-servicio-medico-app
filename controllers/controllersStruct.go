package controllers

import (
	"github.com/vengador20/sistema-servicios-medicos/database"
)

type Controllers struct {
	Client database.DatabaseMongodb //*mongo.Client
}
