package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"comments_service/graph/model"
	"context"
	"log"
)

// Register is the resolver for the Register field.
func (r *mutationResolver) Register(ctx context.Context, registerData model.RegisterData) (*model.RegisterStatus, error) {
	token, err := r.Uc.Register(ctx, registerData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.RegisterStatus{Token: &token}, nil
}

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, identificationData model.IdentificationData, postData string, isCommentEnbale *bool) (*model.CreateStatus, error) {
	if err := r.Uc.CreatePost(ctx, identificationData, postData, isCommentEnbale); err != nil {
		log.Println(err)
		return nil, err
	}
	resp := "Successfully published"
	return &model.CreateStatus{Result: &resp}, nil
}

// AddComment is the resolver for the AddComment field.
func (r *mutationResolver) AddComment(ctx context.Context, identificationData model.IdentificationData, comment model.SComment) (*model.CreateStatus, error) {
	if err := r.Uc.AddComment(identificationData, comment); err != nil {
		log.Println(err)
		return nil, err
	}
	resp := "Successfully published"
	return &model.CreateStatus{Result: &resp}, nil
}

// Posts is the resolver for the Posts field.
func (r *queryResolver) Posts(ctx context.Context, limit *int, offset *int) ([]*model.Post, error) {
	posts, err := r.Uc.Posts(ctx, *limit, *offset)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}

// PostAndComment is the resolver for the PostAndComment field.
func (r *queryResolver) PostAndComment(ctx context.Context, postID *string, limit *int, offset *int) (*model.PostWithComment, error) {
	if PWC, err := r.Uc.PostAndComment(postID, *limit, *offset); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return PWC, nil
	}
}

// GetCommentsFromPost is the resolver for the GetCommentsFromPost field.
func (r *subscriptionResolver) GetCommentsFromPost(ctx context.Context, identificationData model.IdentificationData, postID string) (<-chan *model.RComment, error) {
	return r.Uc.GetCommentsFromPost(ctx, identificationData, postID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
