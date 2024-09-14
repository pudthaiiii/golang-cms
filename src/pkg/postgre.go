package pkg

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go-ibooking/src/app/model"
	log "go-ibooking/src/pkg/logger"
)

func NewPgDatastore() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	ssl := os.Getenv("DB_SSL")

	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", host, port),
		User:   url.UserPassword(user, password),
		Path:   dbName,
	}

	query := u.Query()
	query.Set("sslmode", ssl)
	u.RawQuery = query.Encode()

	dsn := u.String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &sqlLogger{},
	})

	if err != nil {
		log.Log.Err(err).Msg("failed to connect to PostgreSQL")
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Log.Err(err).Msg("failed to get SQL database instance")
		return nil
	}

	// AutoMigrate จะสร้างตารางที่คุณกำหนดในโมเดล
	if err := db.AutoMigrate(&model.Merchant{}, &model.User{}, &model.Role{}); err != nil {
		log.Log.Err(err).Msg("Migration failed")
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	log.Write.Info().Msg("Successfully connected to PostgreSQL")
	return db
}

type sqlLogger struct {
	logger.Interface
}

func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if os.Getenv("DB_LOGGING") == "false" {
		return
	}

	sql, _ := fc()
	log.Write.Printf("SQL: %s\n==============================================================================\n", sql)
}
