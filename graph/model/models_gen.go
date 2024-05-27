// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateStatus struct {
	Result *string `json:"result,omitempty"`
}

type IdentificationData struct {
	Login string `json:"login"`
	Token string `json:"token"`
}

type Mutation struct {
}

type Post struct {
	Author          *string `json:"author,omitempty"`
	TimeAdd         *string `json:"timeAdd,omitempty"`
	Subject         *string `json:"subject,omitempty"`
	ID              *string `json:"ID,omitempty"`
	IsCommentEnbale *bool   `json:"isCommentEnbale,omitempty"`
}

type PostWithComment struct {
	Author          *string     `json:"author,omitempty"`
	TimeAdd         *string     `json:"timeAdd,omitempty"`
	Subject         *string     `json:"subject,omitempty"`
	ID              *string     `json:"ID,omitempty"`
	IsCommentEnbale *bool       `json:"isCommentEnbale,omitempty"`
	Comments        []*RComment `json:"comments,omitempty"`
}

type Query struct {
}

type RComment struct {
	CommentData  *string `json:"CommentData,omitempty"`
	ParentID     *string `json:"ParentID,omitempty"`
	PostID       *string `json:"PostID,omitempty"`
	CommentID    *string `json:"CommentID,omitempty"`
	NestingLevel *int    `json:"NestingLevel,omitempty"`
	TimeAdd      *string `json:"timeAdd,omitempty"`
}

type RegisterData struct {
	Login string `json:"login"`
}

type RegisterStatus struct {
	Token *string `json:"token,omitempty"`
}

type SComment struct {
	CommentData string  `json:"CommentData"`
	ParentID    *string `json:"ParentID,omitempty"`
	PostID      string  `json:"PostID"`
}

type Subscription struct {
}
