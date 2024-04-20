package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	cfg := NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DbName,
		cfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("repository connection error; %w", err)
	}

	return db, nil
}
