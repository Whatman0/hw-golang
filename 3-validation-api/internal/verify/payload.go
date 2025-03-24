package verify

type Mail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address" validate:"required,email"`
}

type ToJson struct {
	Address string `json:"address"`
	Hash    string `json:"hash"`
}
