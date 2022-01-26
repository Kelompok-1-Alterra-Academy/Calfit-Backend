package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateNews struct {
	Title               string `json:"title" form:"title"`
	Description         string `json:"description" form:"description"`
	Content             string `json:"content" form:"content"`
	Operational_adminID int    `json:"operationalAdminId" form:"operationalAdminId"`
}
