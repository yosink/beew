package models

type User struct {
	BaseModel
	Phone       string `json:"phone"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	State       uint8  `json:"state"`
	Password    string `json:"-"`
	Description string `json:"description"`
}

func (u User) TableName() string {
	return "users"
}
