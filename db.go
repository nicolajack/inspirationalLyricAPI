package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// initializes table if database isn't found
func initDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	db, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %w", err)
	}

	// create the lyric table
	_, err = db.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS lyrics (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            artist TEXT NOT NULL,
            lyric TEXT NOT NULL
        )
    `)
	if err != nil {
		return fmt.Errorf("unable to create table: %w", err)
	}

	// seed data only if the table is empty
	var count int
	err = db.QueryRow(context.Background(), "SELECT COUNT(*) FROM lyrics").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		seeds := []lyric{
			{Title: "Ego Death at a Bachelorette Party", Artist: "Hayley Williams", Lyric: "Can only go up from here."},
			{Title: "Sun Bleached Flies", Artist: "Ethel Cain", Lyric: "If it's meant to be then it'll be (oh) / I forgive it all as it comes back to me (back to me)"},
			{Title: "It'll All Work Out", Artist: "Phoebe Bridgers", Lyric: "That's the way it goes, it'll all work out."},
		}
		for _, s := range seeds {
			_, err = db.Exec(context.Background(),
				"INSERT INTO lyrics (title, artist, lyric) VALUES ($1, $2, $3)",
				s.Title, s.Artist, s.Lyric,
			)
			if err != nil {
				return fmt.Errorf("unable to seed data: %w", err)
			}
		}
	}

	return nil
}
