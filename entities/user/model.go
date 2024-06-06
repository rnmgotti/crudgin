package user

type Users struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
