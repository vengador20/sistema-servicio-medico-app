package config

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/vengador20/sistema-servicios-medicos/database/models"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type UserT interface {
	*models.UserRegister | *models.UserLogin | *models.ServicioRegister
}

func ValidateUser[userT UserT](user userT) ([]string, error) {
	//validate := validator.New()

	enLocate := en.New()
	esTrans := es.New()

	//ingles a español
	uni = ut.New(enLocate, enLocate, esTrans)

	//obtener el traductor de español
	trans, _ := uni.GetTranslator("es")

	validate = validator.New()

	//traducir el validator a el traductor
	es_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)

	var message []string

	if err != nil {
		//obtener los errores del validator
		errs := err.(validator.ValidationErrors)

		//fmt.Println(errs.Translate(trans))

		//traduce los errores de la estructura
		for _, v := range errs.Translate(trans) {
			message = append(message, replaceString(v))
		}

		fmt.Printf("message: %v\n", message)
		return message, fmt.Errorf("error")
	}

	return message, nil

}

func replaceString(str string) string {
	spl := strings.Split(str, " ")

	switch spl[0] {
	case "Password":
		return strings.Replace(str, "Password", "La contraseña", 3)

	case "Age":
		return strings.Replace(str, "Age", "La edad", 3)

	case "Apellidos":
		return strings.Replace(str, "Apellidos", "Los Apellidos", 3)

	case "Email":
		return strings.Replace(str, "Email", "El correo electrónico", 3)

	case "Name":
		return strings.Replace(str, "Name", "El Nombre", 3)
	case "Id":
		return strings.Replace(str, "Id", "Selecciona un servicio médico", 3)

	default:
		return str
	}
}

// validar cita medica
func Validate[T any](Tipo T) ([]string, error) {
	//validate := validator.New()

	enLocate := en.New()
	esTrans := es.New()

	//ingles a español
	uni = ut.New(enLocate, enLocate, esTrans)

	//obtener el traductor de español
	trans, _ := uni.GetTranslator("es")

	validate = validator.New()

	//traducir el validator a el traductor
	es_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(Tipo)

	var message []string

	if err != nil {
		//obtener los errores del validator
		errs := err.(validator.ValidationErrors)

		//fmt.Println(errs.Translate(trans))

		//traduce los errores de la estructura
		for _, v := range errs.Translate(trans) {
			message = append(message, replaceString(v))
		}

		fmt.Printf("message: %v\n", message)
		return message, fmt.Errorf("error")
	}

	return message, nil

}
