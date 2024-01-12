package entity

type Video struct {
	//Title       string `json:"title"`
	//Description string `json:"description"`
	//URL         string `json:"url"`
	Items []struct {
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		}
	}
}
