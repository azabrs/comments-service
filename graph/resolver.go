package graph

import commentusercase "comments_service/internal/commentUserCase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	uc commentusercase.UserCase
}
