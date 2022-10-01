package services

import (
	"fmt"
	"time"

	"github.com/rmwebdev/jojomic-restapi/app/migrations"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB Main database connection
var DB *gorm.DB

// InitDatabase initialize database connection
func InitDatabase() {
	if nil == DB {
		db := dbConnect()
		if nil != db {
			DB = db

			dbMigrate()
		}
	}
}

func dbConnect() *gorm.DB {
	logLevel := logger.Info

	switch viper.GetString("ENVIRONMENT") {
	case "staging":
		logLevel = logger.Error
	case "production":
		logLevel = logger.Silent
	}

	config := gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &config)

	if nil != err {
		panic(err)
	}

	if nil != db {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetConnMaxLifetime(time.Second * 5)
	}

	return db
}

func dbMigrate() {
	db := dbConnect()
	if nil != db && len(migrations.ModelMigrations) > 0 {
		err := db.AutoMigrate(migrations.ModelMigrations...)

		if nil != err {
			panic(err)
		}
		db.Migrator().DropTable("schema_migration")

		sqlDB, _ := db.DB()

		defer sqlDB.Close()
	}
}
