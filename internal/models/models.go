package models

import "time"

type User struct{
	Login          string `json:"login" db:"login"`
	HashedPassword string `json:"password_hash" db:"password_hash"`
}

type Post struct{
	Author string `json:"author" db:"author"`
	TimeAdd time.Time `json:"time_add" db:"time_add"`
	IsCommentEnable bool `json:"is_comment_enable" db:"is_comment_enable"`
	Subject string `json:"subject" db:"subject"`
}