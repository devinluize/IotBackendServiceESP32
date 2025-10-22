package configenv

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"net/url"
	"time"

	goredis "github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Database struct {
	Client *goredis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func InitTestDB() *gorm.DB {
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	//	dsn := fmt.Sprintf("sqlserver://localhost:1433?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient", EnvConfigs.DBName)
	//dsn := fmt.Sprintf("sqlserver://%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
	//	EnvConfigs.Hostname,
	//	EnvConfigs.DBPort,
	//	EnvConfigs.DBName)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
		EnvConfigs.DBUser,
		EnvConfigs.DBPass,
		EnvConfigs.DBHost,
		EnvConfigs.DBPort,
		EnvConfigs.DBName)
	//dsn := "sqlserver://localhost:1433?database=assignmentDB&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		log.Fatal("Cannot connected database ", err)
		return nil
	}
	sqlDB, _ := db.DB()

	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("Request Timeout ", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	//log.Info("Connected Database " + EnvConfigs.DBDriver + "- Connected on : " + EnvConfigs.DBHost)

	return db
}
func InitDB() *gorm.DB {
	return &gorm.DB{}
	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")
	//	dsn := fmt.Sprintf("sqlserver://localhost:1433?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient", EnvConfigs.DBName)
	//dsn := fmt.Sprintf("sqlserver://%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
	//	EnvConfigs.Hostname,
	//	EnvConfigs.DBPort,
	//	EnvConfigs.DBName)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient",
		EnvConfigs.DBUser,
		EnvConfigs.DBPass,
		EnvConfigs.DBHost,
		EnvConfigs.DBPort,
		EnvConfigs.DBName)
	//dsn := "sqlserver://localhost:1433?database=assignmentDB&connection+timeout=30&encrypt=disable&trustServerCertificate=false&app name=SqlClient"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "dbo.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		log.Fatal("Cannot connected database ", err)
		return nil
	}
	sqlDB, _ := db.DB()

	err = sqlDB.Ping()

	if err != nil {
		log.Fatal("Request Timeout ", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	log.Info("Connected Database " + EnvConfigs.DBDriver + "- Connected on : " + EnvConfigs.DBHost)

	return db
}

func InitCloudinary() *cloudinary.Cloudinary {
	cld, errr := cloudinary.NewFromURL("cloudinary://695971277991789:jXnWGXSCY230XQ_5QUtMGcb9T18@dlrd9z1mk")
	if errr != nil {
		log.Fatal("Request Timeout ", errr)
	}
	cld.Config.URL.Secure = true
	return cld
}
func InitRedisDB() *Database {
	client := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%v", ""),
		Username: "",
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis Error", err)
		return nil
	}
	log.Info("Connected Redis Database ")

	return &Database{
		Client: client,
	}
}
