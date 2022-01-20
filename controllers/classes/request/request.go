package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateClass struct {
	Name               string `json:"name" form:"name"`
	Description        string `json:"description" form:"description"`
	Banner_picture_url string `json:"bannerPictureUrl" form:"bannerPictureUrl"`
	Card_picture_url   string `json:"cardPictureUrl" form:"cardPictureUrl"`
	Online             bool   `json:"online" form:"online"`
	Link               string `json:"link" form:"link"`
	Category           string `json:"category" form:"category"`
	Status             string `json:"status" form:"status"`
	Price              int    `json:"price" form:"price"`
	Membership_typeID  int    `json:"membershipTypeId" form:"membershipTypeId"`
}
