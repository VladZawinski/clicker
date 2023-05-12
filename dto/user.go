package dto

type User struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ClickedPostUser struct {
	Id    int  `json:"id"`
	User  User `json:"user"`
	Count int  `json:"count"`
}
