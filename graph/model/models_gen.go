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

type Query struct {
}

type RegisterData struct {
	Login string `json:"login"`
}

type RegisterStatus struct {
	Token *string `json:"token,omitempty"`
}
