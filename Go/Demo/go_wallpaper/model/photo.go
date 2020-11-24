package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	PhotoId           int       `json:"id"`
	PhotoCreatedAt    int64		`json:"created_at"`
	Privacy           bool      `json:"privacy"`
	Profile           bool      `json:"profile"`
	URL               string    `json:"url"`
	UserID            int       `json:"user_id"`
	Status            int       `json:"status"`
	Width             int       `json:"width"`
	Height            int       `json:"height"`
	Rating            float64   `json:"rating"`
	HighestRating     float64   `json:"highest_rating"`
	HighestRatingDate int64		`json:"highest_rating_date"`
	ImageFormat       string    `json:"image_format"`
	//Images            []struct {
	//	Format   string `json:"format"`
	//	Size     int    `json:"size"`
	//	URL      string `json:"url"`
	//	HTTPSURL string `json:"https_url"`
	//} `json:"images"`
	ImageURL           string    `json:"image_url"`
	Name               string      `json:"name"`
	//Description        string      `json:"description"`  长文本考虑二进制存储
	Category           int         `json:"category"`
	TakenAt            int64   `json:"taken_at"`
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
	//Liked              interface{} `json:"liked"`
	//Voted              interface{} `json:"voted"`
	CommentsCount      int         `json:"comments_count"`
	VotesCount         int         `json:"votes_count"`
	PositiveVotesCount int         `json:"positive_votes_count"`
	TimesViewed        int         `json:"times_viewed"`

	EditorsChoice     bool        `json:"editors_choice"`
	//EditorsChoiceDate interface{} `json:"editors_choice_date"`
	//EditoredBy        interface{} `json:"editored_by"`
	Feature           string      `json:"feature"`
	FeatureDate       int64   	  `json:"feature_date"`

	//FillSwitch 		  Fillswitch   `json:"fill_switch"`
	//
	//User              User		   `json:"user"`
}

func (photo *Photo) SavePhoto() {

	if !db.HasTable(&Photo{}) {
		//直接通过 db.CreateTable 就可以创建表了，非常方便，
		//还可以通过 db.Set 设置一些额外的表属性,
		fmt.Println("尝试建表 ")

		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&Photo{}).Error; err != nil {
			fmt.Println("建表错误,err", err)
			panic(err)
		}
	}


	createDb := db.Create(&photo)
	err := createDb.Error
	if err != nil {
		fmt.Println("新增照片数据错误,err", err)
		panic(err)
	}
	fmt.Println("插入照片成功")
	fmt.Println(photo.ImageURL)

}

func FindAllPhotos() []Photo {

	var result []Photo

	createDb := db.Find(&result)
	err := createDb.Error
	if err != nil {
		fmt.Println("查询所有photo数据错误,err", err)
		panic(err)
	}
	return result
}