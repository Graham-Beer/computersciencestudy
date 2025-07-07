package api

import (
	"computersciencestudy/internal"
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func SlugHandler(w http.ResponseWriter, r *http.Request) {
	// Extract slug from /api/posts/:slug
	slug := r.URL.Path[len("/api/posts/"):]

	connStr := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var p internal.Post
	err = db.QueryRow("SELECT id, title, content, author, created_at, slug, image_url FROM posts WHERE slug = $1", slug).
		Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug, &p.ImageUrl)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
