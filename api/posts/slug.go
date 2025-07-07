package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
	ImageUrl  string `json:"image_url"` // Extended field
}

func Handler(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/api/post/"):]

	connStr := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var p Post
	err = db.QueryRow("SELECT id, title, content, author, created_at, slug, image_url FROM posts WHERE slug = $1", slug).
		Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug, &p.ImageUrl)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*") // Add CORS for frontend fetch
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
