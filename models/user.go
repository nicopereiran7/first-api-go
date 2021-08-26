package models

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
}
