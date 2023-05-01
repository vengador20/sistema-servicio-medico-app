package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServiciosFuneraria struct {
	Nombre string             `json:"nombre" validate:"required" bson:"nombre"`
	Costo  float64            `json:"costo" validate:"required" bson:"costo"`
	IdUser primitive.ObjectID `json:"idUser" validate:"required" bson:"idUser"`
}

type Cita struct {
	Fecha          string `json:"fecha" validate:"required"`
	Hora           string `json:"hora" validate:"required"`
	NombreCompleto string `json:"pacienteNombre" validate:"required"`
	Telefono       uint64 `json:"telefono" validate:"required"`
	Alergias       string `json:"alergias" validate:"required"`
	IdUser         string `json:"idUser" validate:"required"`
	IdServicio     string `json:"idServicio" validate:"required"`
}

type ServicioCitas struct {
	Fecha          string        `json:"fecha" validate:"required"`
	Hora           string        `json:"hora" validate:"required"`
	NombreCompleto string        `json:"pacienteNombre" validate:"required"`
	Telefono       uint64        `json:"telefono" validate:"required"`
	Alergias       string        `json:"alergias" validate:"required"`
	IdUser         []User        `json:"idUser" validate:"required"`
	IdServicio     []UserService `json:"idServicio" validate:"required"`
}

type ServiciOxigeno struct {
	Nombre     string             `json:"nombre" validate:"required"`
	Costo      float64            `json:"costo" validate:"required"`
	Tipo       string             `json:"tipo" validate:"required"` //renta o venta
	IdServicio primitive.ObjectID `json:"idServicio" validate:"required"`
}

type GetServiciOxigeno struct {
	Nombre     string  `json:"nombre" validate:"required"`
	Costo      float64 `json:"costo" validate:"required"`
	Tipo       string  `json:"tipo" validate:"required"` //renta o venta
	IdServicio primitive.ObjectID/*[]UserService*/ `json:"idServicio" validate:"required"`
}

type Medicamento struct {
	Nombre     string             `json:"nombre" validate:"required"`
	Imagen     string             `json:"imagen" validate:"required"`
	Precio     float64            `json:"precio" validate:"required"`
	Domicilio  bool               `json:"domicilio" validate:"required"`
	IdServicio primitive.ObjectID `json:"idServicio" validate:"required"`
}

type MedicamentoModel struct {
	Nombre     string        `json:"nombre" validate:"required"`
	Imagen     string        `json:"imagen" validate:"required"`
	Precio     float64       `json:"precio" validate:"required"`
	Domicilio  bool          `json:"domicilio" validate:"required"`
	IdServicio []UserService `json:"idServicio" validate:"required"`
}

// (imagen del medicamento, precio,
//dirección de la farmacia, si se hacen entregas a domicilio o no).
