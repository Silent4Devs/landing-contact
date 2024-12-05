package databases

import (
	"fmt"
	"log"

	"fiber-boilerplate/config"
	"fiber-boilerplate/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func BuildDatabaseURI() string {
	dbType := config.GetEnvValue("DBConnection")
	switch dbType {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.GetEnvValue("DBUser"),
			config.GetEnvValue("DBPassword"),
			config.GetEnvValue("DBHost"),
			config.GetEnvValue("DBPort"),
			config.GetEnvValue("DBName"),
		)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			config.GetEnvValue("DBHost"),
			config.GetEnvValue("DBPort"),
			config.GetEnvValue("DBUser"),
			config.GetEnvValue("DBName"),
			config.GetEnvValue("DBPassword"),
		)
	case "sqlite":
		return config.GetEnvValue("DBName")
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
		return ""
	}
}

func Connect() error {
	var err error

	dbURI := BuildDatabaseURI()
	dbConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}

	switch config.GetEnvValue("DBConnection") {
	case "mysql":
		Database, err = gorm.Open(mysql.Open(dbURI), dbConfig)
	case "postgres":
		Database, err = gorm.Open(postgres.Open(dbURI), dbConfig)
	case "sqlite":
		Database, err = gorm.Open(sqlite.Open(dbURI), dbConfig)
	default:
		log.Fatalf("Unsupported database type")
		return fmt.Errorf("unsupported database type")
	}

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	Database.Logger = logger.Default.LogMode(logger.Info)

	// Migrate models
	if err := Database.AutoMigrate(&models.User{}, &models.PasswordReset{}); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	return nil
}
