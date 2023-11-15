package routes

import (
	"apidokter/controllers"
	"apidokter/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	dokter := r.Group("/api")

	dokter.Get("/", middlewares.Auth, controllers.Index)
	dokter.Get("/:kode_dokter", middlewares.Auth, controllers.Show)
	dokter.Post("/", middlewares.Auth, controllers.Create)
	dokter.Put("/:kode_dokter", middlewares.Auth, controllers.Update)
	dokter.Delete("/:kode_dokter", middlewares.Auth, controllers.Delete)
}