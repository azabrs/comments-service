package commentusecase

import (
	"comments_service/graph/model"
	secure_access "comments_service/internal/secureAccess/mocks"
	storage "comments_service/internal/storage/mocks"
	"context"
	"reflect"
	"sync"
	"testing"

)

func TestUseCase_Posts(t *testing.T) {
	type fields struct {
		stor           storage.Storage
		secure         secure_access.SecureAccess
		subCh          []chan *model.RComment
		freSlot        []int
		occupiedSlot   []int
		maxSubscribers int
		m              *sync.RWMutex
		MaxCommentSize int
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Post
		wantErr bool
	}{
		struct{name string; fields fields; args args; want []*model.Post; wantErr bool}{name: "firstCase", args: args{
			limit: 1,
			offset: 0,
		}, want: []*model.Post{
		}},
	}
	for _, tt := range tests {
		stor := storage.NewStorage(t)
		secure := secure_access.NewSecureAccess(t)

		stor.On("Posts", tt.args.limit, tt.args.offset).Return([]*model.Post{
			}, nil)

		t.Run(tt.name, func(t *testing.T) {
			u := &UseCase{
				stor:           stor,
				secure:         secure,
				subCh:          tt.fields.subCh,
				freSlot:        tt.fields.freSlot,
				occupiedSlot:   tt.fields.occupiedSlot,
				maxSubscribers: tt.fields.maxSubscribers,
				m:              tt.fields.m,
				MaxCommentSize: tt.fields.MaxCommentSize,
			}
			got, _ := u.Posts(tt.args.ctx, tt.args.limit, tt.args.offset)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.Posts() = %v, want %v", got, tt.want)
			}
		})
	}
}

