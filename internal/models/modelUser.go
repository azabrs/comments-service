package models
type User struct{
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"password_hash" db:"password_hash"`
}