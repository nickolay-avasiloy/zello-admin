package main

import (
    "log"
    "os"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations() {
    // Load environment variables from .env file
    connStr := os.Getenv("POSTGRES_CONNECTION_STRING")

    m, err := migrate.New(
        "file://migrations",
        connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }

    log.Println("Migrations applied successfully")
}