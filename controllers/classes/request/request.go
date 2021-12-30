package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateClass struct {
	Name               string `json:"name" form:"name"`
	Description        string `json:"description" form:"description"`
	Banner_picture_url string `json:"banner_picture_url" form:"bannerPictureUrl"`
	Card_picture_url   string `json:"card_picture_url" form:"cardPictureUrl"`
	Category           string `json:"category" form:"category"`
	Status             string `json:"status" form:"status"`
}
