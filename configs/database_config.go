package configs

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Database IDatabaseConfig
	once     sync.Once
)

type IDatabaseConfig interface {
	DbUser() *gorm.DB
}

type databaseConfig struct {
	dbUser *gorm.DB
}

func (d *databaseConfig) DbUser() *gorm.DB {
	return d.dbUser
}

var gormConfig = gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
}

func loadConfig() (*databaseConfig, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "103.28.52.104"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASS", "KenjaKu10"),
		getEnv("DB_NAME", "postgres"),
	)

	dbUser, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, err
	}

	return &databaseConfig{
		dbUser: dbUser,
	}, nil
}

func NewDatabaseConfig() IDatabaseConfig {
	configDatabase, err := loadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load database config: %v", err)
	}
	return configDatabase
}

func CheckEnvDatabase() error {
	_, err := loadConfig()
	return err
}

func ReloadDatabaseConfig() {
	configDatabase, err := loadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to reload database config: %v", err)
	}
	Database = configDatabase
	log.Println("✅ Database config reloaded")
}

func init() {
	once.Do(func() {
		Database = NewDatabaseConfig()
	})
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
