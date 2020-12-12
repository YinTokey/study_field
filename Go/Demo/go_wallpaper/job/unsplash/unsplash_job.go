package unsplash

// https://api.unsplash.com/photos/?client_id=YOUR_ACCESS_KEY
var unsplash_base_url string = "https://api.unsplash.com"

type UnsplashJob struct {
}

func NewUnsplashJob() *UnsplashJob {
	return &UnsplashJob{}
}

func (u *UnsplashJob) fetchPics() {

}
