package manager

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var onceLoadDB sync.Once

type InfraManager interface {
	GetDB() *sql.DB
}

type infraManager struct {
	db *sql.DB
}

func (im *infraManager) GetDB() *sql.DB {
	onceLoadDB.Do(func() {
		db, err := sql.Open("postgres", "user=postgres host=localhost password=12345 dbname=laundry sslmode=disable")
		if err != nil {
			log.Fatal("Cannot start app, error when connect to DB", err.Error())
		}

		im.db = db
	})
	return im.db
}

func NewInfraManager() InfraManager {
	return &infraManager{}
}
