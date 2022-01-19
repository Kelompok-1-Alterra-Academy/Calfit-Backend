package request

type KeyStruct struct {
	Key string `json:"key"`
}

type Memberships struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
}
