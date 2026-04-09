package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getLyrics responds with the list of all lyrics as JSON.
func getLyrics(c *gin.Context) {
	rows, err := db.Query(context.Background(), "SELECT title, artist, lyric FROM lyrics")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query lyrics"})
		return
	}
	defer rows.Close()

	var results []lyric
	for rows.Next() {
		var l lyric
		if err := rows.Scan(&l.Title, &l.Artist, &l.Lyric); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan row"})
			return
		}
		results = append(results, l)
	}

	c.IndentedJSON(http.StatusOK, results)
}

// getRandomLyric responds with a random lyric
func getRandomLyric(c *gin.Context) {
	var l lyric
	err := db.QueryRow(context.Background(),
		"SELECT title, artist, lyric FROM lyrics ORDER BY RANDOM() LIMIT 1",
	).Scan(&l.Title, &l.Artist, &l.Lyric)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get random lyric"})
		return
	}

	c.IndentedJSON(http.StatusOK, l)
}
