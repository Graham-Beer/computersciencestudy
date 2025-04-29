package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get slug from URL path
	slug := r.URL.Path[len("/api/posts/"):]

	// Connect to Neon database
	connStr := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Fetch post by slug
	var p Post
	err = db.QueryRow("SELECT id, title, content, author, created_at, slug FROM posts WHERE slug = $1", slug).
		Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
