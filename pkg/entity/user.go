package entity

// User : Database model for user

type User struct {
	ID       string `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
	DOB      string `json:"dob"`
}
