package models

type Role struct {
	Role_id   uint   `json:"role_id" gorm:"primaryKey"`
	Role_name string `json:"role_name"`
}
