package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type SQLiteStore struct {
	addr string 
}

func NewDBConnection(addr string) *SQLiteStore {
	return &SQLiteStore{
		addr: addr,
	}
} 

func (this *SQLiteStore) initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", this.addr)
	if err != nil {
		return nil, err 
	}

	return db, nil 
}

func (this *SQLiteStore) SetupDB() *sql.DB{
	db, err := this.initDB()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	log.Println("Base de datos conecatada exitosamente")
	return db
}
