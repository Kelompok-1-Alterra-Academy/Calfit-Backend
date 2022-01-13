package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateGym struct {
	Name                 string `json:"name" form:"name"`
	Description          string `json:"description" form:"description"`
	Telephone            string `json:"telephone" form:"telephone"`
	Picture              string `json:"picture" form:"picture"`
	Operational_admin_ID uint   `json:"operationalAdminId" form:"operationalAdminId"`
	// Address_ID           uint `json:"address_id" form:"address_id"`
	Address     string `json:"address" form:"address"`
	District    string `json:"district" form:"district"`
	City        string `json:"city" form:"city"`
	Postal_code string `json:"postalCode" form:"postalCode"`
}
