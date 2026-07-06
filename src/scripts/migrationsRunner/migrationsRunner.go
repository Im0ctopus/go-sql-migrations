package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var migrationTableQuery = `
CREATE SCHEMA IF NOT EXISTS migrations;

CREATE TABLE IF NOT EXISTS migrations.migrations (
  id SERIAL PRIMARY KEY,
  "scriptName" TEXT NOT NULL,
  timestamp BIGINT NOT NULL
);
`

func main() {
	ctx := context.Background()

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	migrations, err := os.ReadDir("../migrations/")
	if err != nil {
		log.Fatal(err)
	}

	db, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	_, err = db.Exec(ctx, migrationTableQuery)
	if err != nil {
		fmt.Println(err)
	}

	query := `SELECT "scriptName" from migrations.migrations`
	rows, err := db.Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var appliedMigrations []string
	for rows.Next() {
		var scriptName string
		err := rows.Scan(&scriptName)
		if err != nil {
			log.Fatal(err)
		}
		appliedMigrations = append(appliedMigrations, scriptName)
	}

	query = `INSERT INTO migrations.migrations ("timestamp", "scriptName")
	VALUES ($1, $2)
	`
	for _, migration := range migrations {
		if slices.ContainsFunc(appliedMigrations, func(appliendMigration string) bool { return appliendMigration == migration.Name() }) {
			continue
		}

		migrationScript, err := os.ReadFile("../migrations/" + migration.Name())
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(ctx, string(migrationScript))
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec(ctx, query, time.Now().Format("200601021504"), migration.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Applied script " + migration.Name())
	}

}
