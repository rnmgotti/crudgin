package user

type Users struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Login    string `json:"login"`
	Password string `json:"password"`
}