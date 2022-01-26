package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateNews struct {
	Title               string `json:"title" form:"title"`
	Description         string `json:"description" form:"description"`
	Content             string `json:"content" form:"content"`
	Url_Picture         string `json:"url_picture" form:"url_picture"`
	Operational_adminID int    `json:"operationalAdminId" form:"operationalAdminId"`
}
