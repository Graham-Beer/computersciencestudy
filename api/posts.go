package api

import (
	"computersciencestudy/internal"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	connStr := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, content, author, created_at, slug, image_url FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []internal.Post
	for rows.Next() {
		var p internal.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug, &p.ImageUrl)
		if err != nil {
			http.Error(w, "Failed to scan post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
