package models

type Image struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Color       string `json:"color"`
	BlurHash    string `json:"blur_hash"`
	Likes       int    `json:"likes"`
	LikedByUser bool   `json:"liked_by_user"`
	Description string `json:"description"`
	User        struct {
		ID                string `json:"id"`
		Username          string `json:"username"`
		Name              string `json:"name"`
		FirstName         string `json:"first_name"`
		LastName          string `json:"last_name"`
		InstagramUsername string `json:"instagram_username"`
		TwitterUsername   string `json:"twitter_username"`
		PortfolioURL      string `json:"portfolio_url"`
		ProfileImage      struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		Links struct {
			Self   string `json:"self"`
			HTML   string `json:"html"`
			Photos string `json:"photos"`
			Likes  string `json:"likes"`
		} `json:"links"`
	} `json:"user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	Urls                   struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self     string `json:"self"`
		HTML     string `json:"html"`
		Download string `json:"download"`
	} `json:"links"`
}

type SearchImageQuery struct {
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Results    []Image `json:"results"`
}
