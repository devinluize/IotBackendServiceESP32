package main

import (
	configenv "IotBackend/api/config"
	"IotBackend/api/route"
	migration "IotBackend/migrate"
	"fmt"
	"os"
)

//ini code
//	@title			IoT Blynk Smart Agriculture Backend Service
//	@version		1.0
//	@description	IoT Blynk Smart Agriculture Backend Service
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	IoT Blynk Smart Agriculture Backend Service
//	@contact.url IoT
//	@contact.email	devin@gmail.com

//	@license.name	IoT Blynk Smart Agriculture Backend Service

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @host		localhost:3000
// @BasePath
func main() {
	args := os.Args

	env := ""

	if len(args) > 1 {
		env = args[1]
	}
	configenv.InitEnvConfigs(false, env)

	if env == "migrate" {
		fmt.Println("dasdsa")
		migration.Migrate()

		return

	}
	//configenv.InitEnvConfigs(false, env)
	db := configenv.InitDB()

	cld := configenv.InitCloudinary()

	ds := configenv.EnvConfigs.Hostname
	fmt.Println(ds)
	route.StartRouting(db, cld)
}
