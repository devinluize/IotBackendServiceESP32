package route

import (
	"IotBackend/api/controller/blynkController"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

//func AuthRouter(controller auth.AuthController) chi.Router {
//	r := chi.NewRouter()
//	//router.With(middlewares.RouterMiddleware).Post("/", Finishnotecontroller.FinishReceivingNotesRequestMaster)
//	r.Use(cors.Handler(cors.Options{
//		AllowedOrigins:   []string{"*"},
//		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
//		AllowedHeaders:   []string{"Content-Type", "Authorization"},
//		AllowCredentials: true,
//		MaxAge:           300,
//	}))
//	r.Post("/register", controller.Register)
//	r.Post("/login2", controller.AuthLogin)
//
//	r.With(middleware.RouterMiddleware).Get("/login", controller.AuthLogin)
//
//	return r
//}

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
