package api

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	Slug      string `json:"slug"`
	ImageUrl  string `json:"image_url"`
}
