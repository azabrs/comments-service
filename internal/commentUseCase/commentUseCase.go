package commentusecase

import (
	custom_errors "comments_service/errors"
	"comments_service/graph/model"
	"sync"
	"time"

	"comments_service/internal/models"
	secure_access "comments_service/internal/secureAccess"
	"comments_service/internal/storage"
	"context"
	"log"
)

type UseCase struct{
	stor storage.Storage
	secure secure_access.SecureAccess
	subCh []chan *model.RComment
	freSlot []int
	occupiedSlot []int
	maxSubscribers int
	m sync.RWMutex
	MaxCommentSize int
}
func New(stor storage.Storage, secure secure_access.SecureAccess, maxSubscribers int, MaxCommentSize int) UseCase{
	var subCh []chan *model.RComment
	var temp []int
	for i := 0; i < maxSubscribers; i++{
		ch := make(chan *model.RComment)
		subCh = append(subCh, ch)
		temp = append(temp, i)
	}
	
	return UseCase{stor : stor,
						 secure : secure,
						 occupiedSlot: make([]int, 0),
						 freSlot: temp,
					subCh: subCh,
				maxSubscribers: maxSubscribers,}
}

func (u *UseCase)Register(ctx context.Context, registerData model.RegisterData) (string, error){
	token, err := u.secure.Authorize(registerData.Login)
	if err != nil{
		return "", err
	}

	err = u.stor.Register(registerData.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UseCase) CreatePost(ctx context.Context, identificationData model.IdentificationData, postData string, isCommentEnbale *bool) error{
	if err := u.secure.Authentication(identificationData, u.stor); err != nil{
		return err
	}
	Post := models.Post{
		Author : identificationData.Login,
		TimeAdd: time.Now(),
		Subject: postData,
		IsCommentEnable: *isCommentEnbale,
	}
	if err := u.stor.CreatePost(Post); err != nil{
		return err
	}
	return nil
}

func (u *UseCase) Posts(ctx context.Context, limit int, offset int) ([]*model.Post, error){
	posts, err := u.stor.Posts(limit, offset)
	if err != nil{
		return nil, err
	}
	return posts, nil
}

func (u *UseCase) AddComment(identificationData model.IdentificationData, comment model.SComment) error{
	if err := u.secure.Authentication(identificationData, u.stor); err != nil{
		return err
	}
	if len([]rune(comment.CommentData)) > 2000{
		return custom_errors.ErrReachedMaxLetterSize
	}
	if err := u.stor.AddComment(comment, u.subCh, u.occupiedSlot); err != nil{
		return err
	}
	return nil
}

func (u *UseCase) PostAndComment(postID *string, limit int, offset int) (*model.PostWithComment, error) {
	PWC, err := u.stor.PostAndComment(postID, limit, offset)
	if err != nil{
		return nil, err
	}
	return PWC, nil
}
func (u *UseCase) GetCommentsFromPost(ctx context.Context, identificationData model.IdentificationData, postID string) (<-chan *model.RComment, error) {
	ch := make(chan *model.RComment)
	u.secure.Authentication(identificationData, u.stor)
	u.m.Lock()
	defer u.m.Unlock()
	var chNumber int
	if len(u.freSlot) == 0{
		return nil, custom_errors.ErrReachecMaxSub
	}
	for i, val := range(u.freSlot){
		u.freSlot = append(u.freSlot[:i], u.freSlot[i+1:]...)
		u.occupiedSlot = append(u.occupiedSlot, val)
		break
	}
	chNumber = u.occupiedSlot[len(u.occupiedSlot) - 1]
	go func() {
		defer func(){
			u.m.Lock()
			defer u.m.Unlock()
			for i, val := range(u.occupiedSlot){
				if val == chNumber{
					u.occupiedSlot = append(u.occupiedSlot[:i], u.occupiedSlot[i+1:]...)
					u.freSlot = append(u.freSlot, val)
					break
				}
			}
			
		}()

		for {

			select {
			case <-ctx.Done():
				// Exit on cancellation 
				log.Println("Subscription closed.")
				return
			
			case temp :=<- u.subCh[chNumber]:
				if *temp.PostID == postID{
					ch <- temp
				}
			}

		}
	}()
	return ch, nil
}


