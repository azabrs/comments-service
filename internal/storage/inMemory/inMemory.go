package inmemory

import (
	custom_errors "comments_service/errors"
	"comments_service/graph/model"
	"comments_service/internal/models"
	"log"
	"strconv"
	"sync"
	"time"
)


type inmemory struct{
	Users map[string]bool
	PostID int
	Post map[int]model.PostWithComment
	m sync.RWMutex
}

func InitMemory() *inmemory{
	users := make(map[string]bool)
	post := make(map[int]model.PostWithComment)
	return &inmemory{
		Users: users,
		PostID: 0,
		Post: post,
	}
}

func (im *inmemory) Register(login string) error{
	im.m.Lock()
	defer im.m.Unlock()
	_, ok := im.Users[login]
	if ok{
		return custom_errors.ErrAlreadyRegistered
	}
	im.Users[login] = true
	return nil
}

func (im *inmemory)CreatePost(post models.Post) error{
	im.m.Lock()
	defer im.m.Unlock()
	ind := strconv.Itoa(im.PostID)
	buf_s := post.TimeAdd.String()

	im.Post[im.PostID] = model.PostWithComment{
		Post :&model.Post{
			Author : &post.Author,
			ID : &ind,
			Subject : &post.Subject,
			TimeAdd : &buf_s,
			IsCommentEnable : &post.IsCommentEnable,
		},
	}
	im.PostID += 1
	return nil
}

func (im *inmemory)IsLoginExist(login string) error{
	im.m.Lock()
	defer im.m.Unlock()
	_, ok := im.Users[login]
	if !ok{
		return custom_errors.ErrTokenOrUserInvalid
	}
	return nil
}


func (im *inmemory)Posts(limit int) ([]*model.Post, error){
	im.m.RLock()
	defer im.m.RUnlock()
	var posts []*model.Post
	for key, val := range(im.Post){
		ID := strconv.Itoa(key)
		posts = append(posts, &model.Post{
			Author: val.Post.Author,
			IsCommentEnable: val.Post.IsCommentEnable,
			ID: &ID,
			Subject: val.Post.Subject,
			TimeAdd: val.Post.TimeAdd,
		})
	}
	return posts, nil
}

func (im *inmemory)AddComment(Comment model.SComment, subCh []chan *model.RComment, occupiedSlot []int) error{
	im.m.RLock()
	defer im.m.RUnlock()
	avaliable, err := im.IsCommentable(&Comment.PostID)
	if err != nil{
		
		return err
	}
	if !avaliable{
		return custom_errors.ErrPostNotAvaliableToComment
	}
	CurTime := time.Now().String()
	ind, _ := strconv.Atoi(Comment.PostID)
	if Comment.ParentID == nil{
		
		_, ok := im.Post[ind]
		if !ok{
			return custom_errors.ErrPostIDIncorrect 
		}
		temp_s := strconv.Itoa(0)
		
		temp_i := 0
		temp_s2 := strconv.Itoa(im.Post[ind].Count)
		im.Post[ind] = model.PostWithComment{
			Post :im.Post[ind].Post,
			Comments: append(im.Post[ind].Comments, &model.RComment{
				CommentData: &Comment.CommentData,
				ParentID: &temp_s,
				PostID: &Comment.PostID,
				TimeAdd: &CurTime,
				NestingLevel: &temp_i,
				CommentID: &temp_s2,
			}),
			Count: im.Post[ind].Count + 1,
		}
		for _, val := range(occupiedSlot){
			buf := im.Post[ind].Comments[im.Post[ind].Count - 1]
			subCh[val] <- buf
		}
	}else{
		p, ok := im.Post[ind]
		if !ok{
			log.Println(custom_errors.ErrParentIdIncorrect)
		}
		ind2, _ := strconv.Atoi(*Comment.ParentID)
		if *p.Comments[ind2].PostID != Comment.PostID{
			return custom_errors.ErrPostIDIncorrect 
		}  
		nest_level := *p.Comments[ind2].NestingLevel + 1
		temp_s2 := strconv.Itoa(im.Post[ind].Count)
		/*im.Post[ind].Com[im.CommentID] = model.RComment{
			CommentData: &Comment.CommentData,
			ParentID: Comment.ParentID,
			PostID: &Comment.PostID,
			TimeAdd: &CurTime,
			NestingLevel: p.NestingLevel,
			CommentID: &temp_s2,
			
		}*/
		im.Post[ind] = model.PostWithComment{
			Post :im.Post[ind].Post,
			Comments: append(im.Post[ind].Comments, &model.RComment{
				CommentData: &Comment.CommentData,
				ParentID: Comment.ParentID,
				PostID: &Comment.PostID,
				TimeAdd: &CurTime,
				NestingLevel: &nest_level,
				CommentID: &temp_s2,
			}),
			Count: im.Post[ind].Count + 1,
		}
		for _, val := range(occupiedSlot){
			buf := im.Post[ind].Comments[im.Post[ind].Count - 1]
			subCh[val] <- buf
		}
		
	}
	
	return nil
}

func (im *inmemory)PostAndComment(postID *string, limit *int) (*model.PostWithComment, error){
	im.m.RLock()
	defer im.m.RUnlock()
	ind, err := strconv.Atoi(*postID)
	if err != nil{
		return nil, err
	}
	res, ok := im.Post[ind]
	if !ok{
		return nil, custom_errors.ErrPostIDIncorrect
	}
	return &res, nil

}


func (im *inmemory) IsCommentable(postID *string,)(bool, error){
	im.m.RLock()
	defer im.m.RUnlock()
	id, err := strconv.Atoi(*postID)
	if err != nil{
		return false, err
	}
	return *im.Post[id].Post.IsCommentEnable, nil

}