package internalsql

import (
	"EffectiveTask/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgreSQL(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}

	return db, nil
}
