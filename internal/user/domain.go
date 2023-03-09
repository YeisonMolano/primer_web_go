package user

import "time"

type User struct {
	ID        string     `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	FirstName string     `json:"firstName" gorm:"type:char(50);not null"`
	LastName  string     `json:"lastName" gorm:"type:char(50);not null"`
	Email     string     `json:"email" gorm:"type:char(50);not null"`
	Telefono  string     `json:"telefono" gorm:"type:char(10);not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
