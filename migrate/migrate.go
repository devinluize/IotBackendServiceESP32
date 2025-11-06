package migrate

import (
	configenv "IotBackend/api/config"
	"IotBackend/api/entity"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Migrate() {
	configenv.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating to database"
	//dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
	//	configenv.EnvConfigs.DBUser,
	//	configenv.EnvConfigs.DBPass,
	//	configenv.EnvConfigs.DBHost,
	//	configenv.EnvConfigs.DBPort,
	//	configenv.EnvConfigs.DBName)
	//dsn := fmt.Sprintf(
	//	"server=%s\\%s;user id=%s;password=%s;database=%s;encrypt=disable;trustServerCertificate=true",
	//	"34.101.167.83", // host only
	//	"SQLEXPRESS01",  // instance name (no slash in host!)
	//	"sqlserver",
	//	"password",
	//	"smarthome_iot",
	//)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
		configenv.EnvConfigs.DBUser,
		configenv.EnvConfigs.DBPass,
		configenv.EnvConfigs.DBHost,
		configenv.EnvConfigs.DBPort,
		configenv.EnvConfigs.DBName)
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Printf("%s Failed to connect to database with error: %s", logEntry, err)
		panic(err)
	}
	// Get the list of all tables
	fmt.Println("check point")
	var tableNames []string
	db.Raw("SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE'").Scan(&tableNames)
	err = db.AutoMigrate(
		//&entities.EquipmentDifficultyEntities{},
		&entity.BlynkData{},
	)

	if err != nil {
		log.Printf("%s Failed with error: %s", logEntry, err)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
