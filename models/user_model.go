package models

// User model.
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"-"`
}

// TableName give table name of model user.
func (u *User) TableName() string {
	return "users"
}
