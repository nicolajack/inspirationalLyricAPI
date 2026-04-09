package main

import (
	"log"

	"os"

	"github.com/gin-gonic/gin"
)

// lyric represents data about a lyric.
type lyric struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Lyric  string `json:"lyric"`
}

// lyrics slice to seed lyric data.
var lyrics = []lyric{
	{Title: "Ego Death at a Bachelorette Party", Artist: "Hayley Williams", Lyric: "Can only go up from here."},
	{Title: "Sun Bleached Flies", Artist: "Ethel Cain", Lyric: "If it's meant to be then it'll be (oh) / I forgive it all as it comes back to me (back to me)"},
	{Title: "It'll All Work Out", Artist: "Phoebe Bridgers", Lyric: "That's the way it goes, it'll all work out."},
}

func main() {
	// initDB connects to postgresql db
	if err := initDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/lyrics", getLyrics)
	router.GET("/lyrics/random", getRandomLyric)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
}
