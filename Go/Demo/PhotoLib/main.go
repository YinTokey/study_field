package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
	. "PhotoLib/example/gopl"
	//_ "PhotoLib/internal/app"
)


var pupular_url = "https://api.500px.com/v1/photos?feature=popular"

func main () {
	Runex()
	//resp, err := http.Get(pupular_url)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//s , err := ioutil.ReadAll(resp.Body)
	//jsonStr := string(s)
	//
	//var info PxPhotoInfo
	//json.Unmarshal([]byte(jsonStr), &info)
	//
	////fmt.Println(info.Feature)
	//redis.set

}

func Runex() {
	re := Sum(1,2,3,4)
	fmt.Println(re)
}

var redisDB *redis.Client
func InitRedis(){
	//地址，密码
	host := "localhost:6379"
	passwd := "" //没有密码
	redisDB = redis.NewClient(&redis.Options{
		Addr: host,
		Password: passwd,
		DB: 0, //使用默认 DB
	})
}


type PxPhotoInfo struct {
	CurrentPage int    `json:"current_page"`
	TotalPages  int    `json:"total_pages"`
	TotalItems  int    `json:"total_items"`
	Feature     string `json:"feature"`
	Filters     struct {
	} `json:"filters"`
	Photos []struct {
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
		User               struct {
			ID               int       `json:"id"`
			Username         string    `json:"username"`
			Fullname         string    `json:"fullname"`
			AvatarVersion    int       `json:"avatar_version"`
			RegistrationDate time.Time `json:"registration_date"`
			Avatars          struct {
				Tiny struct {
					HTTPS string `json:"https"`
				} `json:"tiny"`
				Small struct {
					HTTPS string `json:"https"`
				} `json:"small"`
				Large struct {
					HTTPS string `json:"https"`
				} `json:"large"`
				Default struct {
					HTTPS string `json:"https"`
				} `json:"default"`
			} `json:"avatars"`
			UserpicURL      string `json:"userpic_url"`
			UserpicHTTPSURL string `json:"userpic_https_url"`
			Usertype        int    `json:"usertype"`
			Active          int    `json:"active"`
			Firstname       string `json:"firstname"`
			Lastname        string `json:"lastname"`
			About           string `json:"about"`
			City            string `json:"city"`
			State           string `json:"state"`
			Country         string `json:"country"`
			CoverURL        string `json:"cover_url"`
			UpgradeStatus   int    `json:"upgrade_status"`
			Affection       int    `json:"affection"`
			FollowersCount  int    `json:"followers_count"`
			Following       bool   `json:"following"`
		} `json:"user"`
		EditorsChoice     bool        `json:"editors_choice"`
		EditorsChoiceDate interface{} `json:"editors_choice_date"`
		EditoredBy        interface{} `json:"editored_by"`
		Feature           string      `json:"feature"`
		FeatureDate       time.Time   `json:"feature_date"`
		FillSwitch        struct {
			AccessDeleted        bool        `json:"access_deleted"`
			AccessPrivate        bool        `json:"access_private"`
			IncludeDeleted       bool        `json:"include_deleted"`
			ExcludePrivate       bool        `json:"exclude_private"`
			ExcludeNude          bool        `json:"exclude_nude"`
			AlwaysExcludeNude    bool        `json:"always_exclude_nude"`
			ExcludeBlock         bool        `json:"exclude_block"`
			CurrentUserID        interface{} `json:"current_user_id"`
			OnlyUserActive       bool        `json:"only_user_active"`
			IncludeTags          bool        `json:"include_tags"`
			IncludeGeo           bool        `json:"include_geo"`
			IncludeLicensing     bool        `json:"include_licensing"`
			IncludeAdminLocks    bool        `json:"include_admin_locks"`
			IncludeLikeBy        bool        `json:"include_like_by"`
			IncludeComments      bool        `json:"include_comments"`
			IncludeUserInfo      bool        `json:"include_user_info"`
			IncludeFollowInfo    bool        `json:"include_follow_info"`
			IncludeEquipmentInfo bool        `json:"include_equipment_info"`
		} `json:"fill_switch"`
	} `json:"photos"`
}