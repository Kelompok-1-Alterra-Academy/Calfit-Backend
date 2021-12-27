package request

type KeyStruct struct {
	Key string `json:"key"`
}

type CreateGym struct {
	Name                 string `json:"name" form:"name"`
	Telephone            string `json:"telephone" form:"telephone"`
	Picture              string `json:"picture" form:"picture"`
	Operational_admin_ID uint   `json:"operational_admin_id" form:"operationalAdminId"`
	// Address_ID           uint `json:"address_id" form:"addressId"`
	Address uint `json:"address_id" form:"address"`
}