package migrations

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"log"
	"os"
)

// RunMigrations execute the migrations
func RunMigrations(dbInstance *gorm.DB, migrations []*gormigrate.Migration) {
	if os.Getenv("ENVIRONMENT") == "" {
		panic("Env ENVIRONMENT is not defined")
	}
	if err := dbInstance.DB().Ping(); err != nil {
		log.Fatalf("Db error: %v", err)
	}
	m := gormigrate.New(dbInstance, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}

// ExecRaw simplify error handling for sql commands
func ExecRaw(db *gorm.DB, command string) bool {
	err := db.Exec(command).Error
	if err != nil {
		message := fmt.Sprintf("Migration Error into %v environment while evaluating command:\n\n ==> %v \n\nERROR:\n\n ==> %v", os.Getenv("ENVIRONMENT"), command, err.Error())
		log.Panicf(message)
		return false
	}
	return true
}
