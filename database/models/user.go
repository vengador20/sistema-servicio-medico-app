package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Nombres  string `json:"nombres" validate:"required"`
	Fecha    string `json:"fecha" validate:"required,datetime"`
	Telefono uint64 `json:"telefono" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UserService struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	//User
	//Nombres       string  `json:"nombres" validate:"required"`
	//NombreUsuario string  `json:"nombreUsuario" validate:"required"`
	NombreNegocio  string `json:"nombreNegocio" validate:"required"`
	NombreCompleto string `json:"nombreCompleto" validate:"required"`

	Telefono uint64  `json:"telefono" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Tipo     uint8   `json:"tipo" bson:"tipo,omitempty" validate:"required"` // 0 es un paciente y 1 es servicio
	Servicio string  `json:"servicio" validate:"required"`
	Latitud  float64 `json:"latitud" validate:"required"`
	Longitud float64 `json:"longitud" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Tipo     uint8  `json:"tipo"`
}

type ServicioRegister struct {
	//Nombres       string  `json:"nombres" validate:"required"`
	//NombreUsuario string  `json:"nombreUsuario" validate:"required"`
	NombreNegocio  string `json:"nombreNegocio" validate:"required"`
	NombreCompleto string `json:"nombreCompleto" validate:"required"`

	Telefono uint64  `json:"telefono" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Tipo     uint8   `json:"tipo" bson:"tipo,omitempty" validate:"required"` // 0 es un paciente y 1 es servicio
	Servicio string  `json:"servicio" validate:"required"`
	Latitud  float64 `json:"latitud" validate:"required"`
	Longitud float64 `json:"longitud" validate:"required"`
	//User
	Cedula   string `json:"cedula" validate:"omitempty,required"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Nombres         string `json:"nombres" bson:"nombres,omitempty" validate:"required"`
	FechaNacimiento string `json:"fecha" bson:"fecha,omitempty" validate:"required"`
	Email           string `json:"email" bson:"email,omitempty" validate:"required,email"`
	Telefono        uint64   `json:"telefono" bson:"telefono,omitempty" validate:"required"` //numeros enteros
	Password        string `json:"password" bson:"password,omitempty" validate:"required"`
	Tipo            uint8  `json:"tipo" bson:"tipo,omitempty"` // 0 es un paciente y 1 es servicio
}

type UserNewPassword struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
