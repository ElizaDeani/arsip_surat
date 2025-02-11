package entity

type User struct {
	Id       int    `json:"id" gorm:"autoIncrement;"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "users"
}
