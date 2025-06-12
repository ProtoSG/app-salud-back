package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ProtoSG/app-salud-back/internal/db/migrations" // tus migraciones .sql y .go
	_ "github.com/lib/pq"                                        // driver de Postgres
	"github.com/pressly/goose/v3"                                // goose v3
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Uso: %s <comando> <cadena_conexión> [args]\n", os.Args[0])
		os.Exit(1)
	}

	cmd := os.Args[1]      // e.g. "up", "down", "status", "down-to"
	dbString := os.Args[2] // e.g. "postgres://user:pass@host:port/dbname?sslmode=disable"
	args := os.Args[3:]    // argumentos adicionales para algunos comandos

	// 1) Abrir conexión
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	// 2) Indicar el dialecto
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose.SetDialect: %v", err)
	}

	// 3) Ejecutar el comando
	//    Firma: Run(cmd string, db *sql.DB, dir string, args ...string) error
	migrationsDir := "internal/db/migrations"
	if err := goose.Run(cmd, db, migrationsDir, args...); err != nil {
		log.Fatalf("goose.Run: %v", err)
	}
}
