package dto

type Login struct {
	Phone    string `json:"username"`
	Password string `json:"password"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	Phone    string `json:"username"`
	Password string `json:"password"`
}
