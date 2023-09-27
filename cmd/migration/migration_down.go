package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/MrRytis/go-weather/internal/migrations"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	color.Green("Running migration down")

	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for input
	color.White("Enter your migration version to rollback to: ")

	// Read the user's input
	input, err := reader.ReadString('\n')
	if err != nil {
		color.Red("Error reading input:", err)
		return
	}

	// Remove spaces and new lines from the input
	input = strings.ReplaceAll(input, " ", "_")
	input = strings.ReplaceAll(input, "\n", "")

	version, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	var db *sql.DB
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open: ", err)
	}
	defer db.Close()

	goose.SetBaseFS(migrations.EmbedMigrationsFs)

	if err = goose.SetDialect("postgres"); err != nil {
		log.Fatal("Failed to set dialect: ", err)
	}

	if err = goose.DownTo(db, "migrations", version); err != nil {
		log.Fatal("Failed migration: ", err)
	}

	color.Green("Migration ran successfully")
}
