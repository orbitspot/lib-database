package database

import (
	"github.com/jinzhu/gorm"
	tools "github.com/orbitspot/lib-database/pkg/gorm"
	"sync"
)

type PostgresRepository struct {
	dbPostgres *gorm.DB
	once       sync.Once
}

type PostgresConfiguration struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

// MySQL var to use
var Postgres = &PostgresRepository{}

// StartPostgres start the DB
func (r *PostgresRepository) Start(configuration PostgresConfiguration) {
	_ = tools.LoadGormPostGres(
		configuration.DBUser,
		configuration.DBPass,
		configuration.DBHost,
		configuration.DBPort,
		configuration.DBName,
		false)
}

// StopPostgres stop the DB
func (r *PostgresRepository) Stop() {
	defer r.dbPostgres.Close()
}

// GetInstance returns a unique instance of gorm.DB
func (r *PostgresRepository) GetInstance() *gorm.DB {
	r.once.Do(func() {
		var err error
		r.dbPostgres, err = tools.GetGormDb()
		if err != nil {
			panic(err.Error())
		}
		r.dbPostgres.SingularTable(true)
		r.dbPostgres.LogMode(true)
	})
	return r.dbPostgres
}
