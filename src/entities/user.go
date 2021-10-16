package entities

type User struct {
	ID              string `json:"id"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Users []*User
