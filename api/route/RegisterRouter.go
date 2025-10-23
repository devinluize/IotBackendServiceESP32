package route

import (
	"IotBackend/api/controller/blynkController"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func BlynkRouter(controller blynkController.BlynkController) chi.Router {
	r := chi.NewRouter()
	//r.Use(middleware.RouterMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Post("/", controller.SendDataToBlynk)

	return r
}
