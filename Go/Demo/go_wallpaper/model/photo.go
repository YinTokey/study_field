package model

import "time"

type Photo struct {
	ID                int       `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Privacy           bool      `json:"privacy"`
	Profile           bool      `json:"profile"`
	URL               string    `json:"url"`
	UserID            int       `json:"user_id"`
	Status            int       `json:"status"`
	Width             int       `json:"width"`
	Height            int       `json:"height"`
	Rating            float64   `json:"rating"`
	HighestRating     float64   `json:"highest_rating"`
	HighestRatingDate time.Time `json:"highest_rating_date"`
	ImageFormat       string    `json:"image_format"`
	Images            []struct {
		Format   string `json:"format"`
		Size     int    `json:"size"`
		URL      string `json:"url"`
		HTTPSURL string `json:"https_url"`
	} `json:"images"`
	ImageURL           []string    `json:"image_url"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Category           int         `json:"category"`
	TakenAt            time.Time   `json:"taken_at"`
	ShutterSpeed       string      `json:"shutter_speed"`
	FocalLength        string      `json:"focal_length"`
	Aperture           string      `json:"aperture"`
	Camera             string      `json:"camera"`
	Lens               string      `json:"lens"`
	Iso                string      `json:"iso"`
	Location           string      `json:"location"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	Nsfw               bool        `json:"nsfw"`
	PrivacyLevel       int         `json:"privacy_level"`
	Watermark          bool        `json:"watermark"`
	HasNsfwTags        bool        `json:"has_nsfw_tags"`
	Liked              interface{} `json:"liked"`
	Voted              interface{} `json:"voted"`
	CommentsCount      int         `json:"comments_count"`
	VotesCount         int         `json:"votes_count"`
	PositiveVotesCount int         `json:"positive_votes_count"`
	TimesViewed        int         `json:"times_viewed"`

	EditorsChoice     bool        `json:"editors_choice"`
	EditorsChoiceDate interface{} `json:"editors_choice_date"`
	EditoredBy        interface{} `json:"editored_by"`
	Feature           string      `json:"feature"`
	FeatureDate       time.Time   `json:"feature_date"`

	FillSwitch 		  Fillswitch   `json:"fill_switch"`

	User              User		   `json:"user"`
}