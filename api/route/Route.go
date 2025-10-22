package route

import (
	EquipmentController "IotBackend/api/controller/Equipment"
	auth3 "IotBackend/api/controller/auth/authImpl"
	menucontroller "IotBackend/api/controller/menu"
	repositoriesEquipmentImpl "IotBackend/api/repositories/Equipment/repositories-equipment-impl"
	MenuImplRepositories "IotBackend/api/repositories/menu/repositories-menu-impl"
	"IotBackend/api/repositories/user/UserRepositoryImpl"
	"IotBackend/api/service/EquipmentService/EquipmentServiceImpl"
	auth2 "IotBackend/api/service/auth"
	menuserviceimpl "IotBackend/api/service/menu/menu-service-impl"
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

	authRepository := UserRepositoryImpl.NewAuthRepoImpl()
	authService := auth2.NewAuthServiceImpl(db, authRepository)
	authController := auth3.NewAuthController(authService)

	ArticleRepository := MenuImplRepositories.NewArticleMenu()
	ArticleService := menuserviceimpl.NewArticleServiceImpl(ArticleRepository, db, cld)
	ArticleController := menucontroller.NewInformatioControllerImpl(ArticleService)

	ProfileRepository := MenuImplRepositories.NewProfileMenuRepositoryImpl()
	ProfileService := menuserviceimpl.NewProfileServiceImpl(db, ProfileRepository)
	ProfileController := menucontroller.NewProfileControllerImpl(ProfileService)

	WeightRepository := MenuImplRepositories.NewWeightHistoryRepositoryImpl()
	WeightService := menuserviceimpl.NewWeightHistoryServiceImpl(db, WeightRepository)
	WeightController := menucontroller.NewWeightHistoryController(WeightService)

	//calendar
	CalendarRepository := MenuImplRepositories.NewEventRepositoryImpl()
	CalendarService := menuserviceimpl.NewCalendarServiceImpl(CalendarRepository, db)
	CalendarController := menucontroller.NewCalendarController(CalendarService)

	//Timer
	TimerRepository := MenuImplRepositories.NewTimerRepositoryImpl()
	TimerService := menuserviceimpl.NewTimerServiceImpl(TimerRepository, db)
	TimerController := menucontroller.NewTimerControllerImpl(TimerService)

	//bookmark
	BookmarkRepository := MenuImplRepositories.NewBookmarkRepositoryImpl()
	BookmarkService := menuserviceimpl.NewBookmarkServiceImpl(db, BookmarkRepository, cld)
	BookmarkController := menucontroller.NewBookmarkController(BookmarkService)

	//equipment course
	EquipmentCourseRepository := repositoriesEquipmentImpl.NewEquipmentCourseRepositoryImpl()
	EquipmentCourseService := EquipmentServiceImpl.NewEquipmentCourseServiceImpl(db, EquipmentCourseRepository, cld)
	EquipmentCourseController := EquipmentController.NewEquipmentCourseControllerImpl(EquipmentCourseService)

	//equipment bookmark
	EquipmentBookmarkRepository := repositoriesEquipmentImpl.NewEquipmentBookmarkRepositoryImpl()
	EquipmentBookmarkService := EquipmentServiceImpl.NewEquipmentBookmarkServiceImpl(EquipmentBookmarkRepository, db, cld)
	EquipmentBookmarkController := EquipmentController.NewEquipmentBookmarkControllerImpl(EquipmentBookmarkService)

	//
	AuthRouter := AuthRouter(authController)
	ArticleRouter := ArticleRouter(ArticleController)
	ProfileRouter := ProfileRouter(ProfileController)
	WeightRouter := WeightRouter(WeightController)
	CalendarRouter := CalendarRouter(CalendarController)
	TimerRouter := TimerRoute(TimerController)
	BookmarkRouter := BookmarkRoute(BookmarkController)
	EquipmentCourseRouter := EquipmentCourseRoute(EquipmentCourseController)
	EquipmentMasterRoute := EquipmentMasterRoute(EquipmentCourseController)
	EquipmentBookmarkRouter := EquipmentBookmarkRoute(EquipmentBookmarkController)
	////////////////////////////////////////////

	router.Mount("/user", AuthRouter)
	router.Mount("/article", ArticleRouter)
	router.Mount("/profile", ProfileRouter)
	router.Mount("/weight", WeightRouter)
	router.Mount("/event", CalendarRouter)
	router.Mount("/timer", TimerRouter)
	router.Mount("/bookmark", BookmarkRouter)
	router.Mount("/equipment/course", EquipmentCourseRouter)
	router.Mount("/equipment", EquipmentMasterRoute)
	router.Mount("/equipment/bookmark", EquipmentBookmarkRouter)
	return router
}
