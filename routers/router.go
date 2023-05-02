package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vengador20/sistema-servicios-medicos/controllers"
	"github.com/vengador20/sistema-servicios-medicos/database"
	"github.com/vengador20/sistema-servicios-medicos/middleware"
)

type Router struct {
	Db database.DatabaseMongodb
}

const (
	funeraria      string = "/servicio-funeraria"
	enfermeros     string = "/enfermeros"
	medico         string = "/medico"
	cita           string = "/cita-medica"
	servicioMedico string = "/perfil/servicio-medico"
	centroSalud    string = "/centro-salud"
	oxigeno        string = "/oxigeno"
	medicamento    string = "/medicamento"
)

func (r *Router) Router(router fiber.Router) {

	controller := controllers.Controllers{
		Client: r.Db,
	}

	//user
	router.Get("/user/:email", controller.GetUserEmail)

	//usuario registro
	router.Post("/login", controller.Login)

	router.Post("/register", controller.RegisterUser)

	//prueba
	router.Delete(funeraria+"/:id", controller.EliminarServicioFuneraria)

	// enefermeros
	router.Get(enfermeros, controller.GetEnfermeros)

	router.Get(enfermeros+"/:id", controller.GetEnfermerosById)

	router.Get(funeraria, controller.GetFuneraria)

	router.Get(funeraria+"/:email", controller.GetFunerariaEmail)

	router.Get(funeraria+"/:id", controller.GetFunerariaById)

	//fmt.Println(funeraria + "/update/:id")

	router.Get(funeraria+"/update-one/:id", controller.GetFunerariaByIdUpdate)

	router.Post(funeraria, controller.CrearServicioFuneraria)

	router.Get(medico, controller.GetMedicos)

	router.Get(medico+"/:id", controller.GetMedicoById)

	router.Get(oxigeno, controller.GetOxigeno)

	router.Get(oxigeno+"/:email", controller.GetOxigenoEmail)

	router.Get(oxigeno+"/:id", controller.GetOxigenoByid)

	router.Post(oxigeno, controller.CrearOxigeno)

	router.Put(oxigeno+"/:id", controller.ModificarOxigeno)

	router.Delete(oxigeno+"/:id", controller.EliminarOxigeno)

	router.Get("/perfil-oxigeno", controller.PerfilOxigeno)

	router.Get(medicamento+"/:name", controller.ConsultaMedicamento)

	router.Get(medicamento, controller.GetMedicamentos)

	router.Post(medicamento, controller.CrearMedicamento)

	router.Get(servicioMedico+"/:email", controller.GetCitas)

	router.Get(cita+"/:id", controller.GetCitaMedicaById)

	router.Put(cita+"/:id", controller.ModificarCita)

	//prueba

	//router.Get("/citas", controller.GetCitas)

	api := router.Group("/servicios")

	//utilizar middleware personalizado
	//valida si el jwt no es modificado
	api.Use("/", middleware.ValidateJwt)

	// funeraria
	api.Get(funeraria, controller.GetFuneraria)

	api.Post(funeraria, controller.CrearServicioFuneraria)

	api.Put(funeraria, controller.ModificarServicioFuneraria)

	api.Delete(funeraria+"/:id", controller.EliminarServicioFuneraria)

	// medicos
	api.Get(medico, controller.GetMedicos)

	api.Get(medico+"/:id", controller.GetMedicoById)

	// enefermeros
	api.Get(enfermeros, controller.GetEnfermeros)

	api.Get(enfermeros+"/:id", controller.GetEnfermerosById)

	// cita medica
	api.Post(cita, controller.AgendarCitaMedica)

	//servicio medico
	api.Get(servicioMedico, controller.GetCitas)

	//centros de salud
	api.Get(centroSalud, controller.GetCentroSalud)

	//oxigeno
	api.Get(oxigeno, controller.GetOxigeno)

	api.Get(oxigeno+"/:id", controller.GetOxigenoByid)

	api.Post(oxigeno, controller.CrearOxigeno)

	api.Put(oxigeno+"/:id", controller.ModificarOxigeno)

}
