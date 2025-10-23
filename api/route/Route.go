package route

import (
	blynkController2 "IotBackend/api/controller/blynkController"
	blynkrepositoryimpl "IotBackend/api/repositories/blynk/blynk-repository-impl"
	"IotBackend/api/service/blynk/BlynkServiceImpl"
	_ "IotBackend/docs"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"gorm.io/gorm"
	"net/http"
)

func StartRouting(db *gorm.DB, cld *cloudinary.Cloudinary) {
	r := chi.NewRouter()
	r.Mount("/api", versionedRouterV1(db, cld))
	swaggerURL := "http://localhost:3000/swagger/doc.json"
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL),
	))
	http.ListenAndServe(":3000", r)

}
func versionedRouterV1(db *gorm.DB, cld *cloudinary.Cloudinary) chi.Router {
	router := chi.NewRouter()

	router.Get("/dev", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write(byte[])
		html := `
	<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Response Page</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f0f0f0;
						margin: 0;
						padding: 20px;
					}
					h1 {
						color: #333;
					}
					p {
						color: #666;
					}
				</style>
			</head>
			<body>
				<h1>Hello, client!</h1>
				<p>This is a response from your Go server.</p>
			</body>
			</html>
    `
		_, err := writer.Write([]byte(html))
		if err != nil {
			http.Error(writer, "Error writing response", http.StatusInternalServerError)
			return
		}
	})
	blynkRepository := blynkrepositoryimpl.NewBlynkRepositoryImpl()
	blynkService := BlynkServiceImpl.NewBlynkServiceImpl(db, blynkRepository)
	blynkController := blynkController2.NewBlynkControllerImpl(blynkService)
	//
	//authRepository := UserRepositoryImpl.NewAuthRepoImpl()
	//authService := auth2.NewAuthServiceImpl(db, authRepository)
	//authController := auth3.NewAuthController(authService)
	blynkRouter := BlynkRouter(blynkController)
	//
	//AuthRouter := AuthRouter(authController)
	////////////////////////////////////////////
	router.Mount("/blynk", blynkRouter)
	//router.Mount("/user", AuthRouter)
	return router
}
