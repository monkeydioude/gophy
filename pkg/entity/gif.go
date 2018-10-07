package entity

type Gif struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	URL        string `json:"url"`
	BitlyURL   string `json:"bitly_url"`
	EmbedURL   string `json:"embed_url"`
	Username   string `json:"username"`
	Source     string `json:"source"`
	Rating     string `json:"rating"`
	ContentURL string `json:"content_url"`
	Title      string `json:"title"`
}

type Gifs struct {
	Data []Gif `json:"data"`
}
