package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
}

var db *sql.DB

func main() {
	// Load .env file
    	err := godotenv.Load()
    	if err != nil {
        	log.Println("No .env file found")
    	}

	// Connect to Neon database
	connStr := os.Getenv("POSTGRES_URL")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/api/posts", getPosts).Methods("GET")
	r.HandleFunc("/api/posts/{slug}", getPost).Methods("GET")

	// Serve static files (HTML, CSS, JS)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, content, author, created_at, slug FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug)
		if err != nil {
			http.Error(w, "Failed to scan post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var p Post
	err := db.QueryRow("SELECT id, title, content, author, created_at, slug FROM posts WHERE slug = $1", slug).
		Scan(&p.ID, &p.Title, &p.Content, &p.Author, &p.CreatedAt, &p.Slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
