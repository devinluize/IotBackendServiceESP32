package migrate

import (
	configenv "IotBackend/api/config"
	entities "IotBackend/api/entities/Equipment"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func Migrate() {
	configenv.InitEnvConfigs(false, "")
	logEntry := "Auto Migrating to database"
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
		&entities.EquipmentDifficultyEntities{},
		&entities.EquipmentTypeEntity{},
		&entities.ForceTypeEntities{},
		&entities.MuscleGroupEntities{},
		&entities.EquipmentCourseDataEntity{},
	)

	if err != nil {
		log.Printf("%s Failed with error: %s", logEntry, err)
		panic(err)
	}

	log.Printf("%s Success", logEntry)
}
