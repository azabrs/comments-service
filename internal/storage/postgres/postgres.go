package postgres

import (
	"comments_service/config"
	custom_errors "comments_service/errors"
	"comments_service/graph/model"
	"comments_service/internal/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct{
	Db *sql.DB
}

func InitDb(c config.PostgresConfig) (*Postgres, error){
	str := fmt.Sprintf("postgres://%v:%v@postgres:%v/%v?sslmode=disable",
		c.User,
		c.Password,
		c.Port,
		c.DbName,
	)
	db, err := sql.Open("postgres", str)
	if err != nil{
		return nil,  fmt.Errorf("cant connect to db, %v", err)
	}
	return &Postgres{Db: db}, nil
}

func (pg *Postgres) Register(login string) error{
	query := `INSERT INTO users(login) VALUES($1)`
	_, err := pg.Db.Exec(query, login)
	if err != nil{
		return custom_errors.ErrAlreadyRegistered
	}


	return nil
}

func (pg *Postgres)CreatePost(post models.Post) error{
	query := `INSERT INTO posts(author, time_add, is_comment_enable, subject) VALUES($1, $2, $3, $4)`
	_, err := pg.Db.Exec(query, post.Author, post.TimeAdd, post.IsCommentEnable, post.Subject)
	if err != nil{
		return err
	}
	return nil
}

func (pg *Postgres)IsLoginExist(login string) error{
	query := `SELECT * FROM users WHERE login = $1`
	s := ""
	d := 0
	if err := pg.Db.QueryRow(query, login).Scan(&d, &s); err != nil{
		return err
	}
	return nil
}

func (pg *Postgres)Posts(limit int, offset int) ([]*model.Post, error){
	var res []*model.Post
	query := `SELECT * FROM posts LIMIT $1 OFFSET $2`
	rows, err := pg.Db.Query(query, limit, offset)
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		var temp model.Post
		err = rows.Scan(&temp.ID, &temp.Author, &temp.TimeAdd, &temp.IsCommentEnable, &temp.Subject)
		if err != nil{
			return nil, err
		}
		res = append(res, &temp)
	}
	return res, nil
}



func (pg *Postgres)AddComment(Comment model.SComment, subCh []chan *model.RComment, occupiedSlot []int) error{
	avaliable, err := pg.IsCommentable(&Comment.PostID)
	if err != nil{
		return err
	}
	if !avaliable{
		return custom_errors.ErrPostNotAvaliableToComment
	}
	query := `INSERT INTO comments(comment_data, parent_id, post_id, time_add, nesting_level) VALUES($1, $2, $3, $4, $5)`
	if Comment.ParentID == nil{
		CurTime := time.Now()
		_, err := pg.Db.Exec(query, Comment.CommentData, 0, Comment.PostID, CurTime, 0)
		if err != nil{
			return err
		}

		TimeAdd := CurTime.String()
		var CommentID string
		queryS2 := `SELECT id FROM comments WHERE time_add = $1`
		if err := pg.Db.QueryRow(queryS2, CurTime).Scan(&CommentID); err != nil{
			return err
		}
		var temp = 0
		var temps = ""
		buf := model.RComment{
			CommentData: &Comment.CommentData,
			ParentID: &temps,
			PostID: &Comment.PostID,
			NestingLevel: &temp,
			TimeAdd: &TimeAdd,
			CommentID: &CommentID,
		}
		for _, val := range(occupiedSlot){
			subCh[val] <- &buf
		}
	}else{
		queryS := `SELECT post_id, nesting_level FROM comments WHERE id = $1`
		var postID string
		var nestingLevel int
		if err := pg.Db.QueryRow(queryS, Comment.ParentID).Scan(&postID, &nestingLevel); err != nil{
			return err
		}
		if postID != Comment.PostID{
			return custom_errors.ErrPostIDIncorrect
		}
		CurTime := time.Now()
		_, err := pg.Db.Exec(query, Comment.CommentData, Comment.ParentID, Comment.PostID, CurTime, nestingLevel + 1)
		if err != nil{
			return err
		}
		TimeAdd := CurTime.String()
		tempLevel := nestingLevel + 1
		var CommentID string
		queryS2 := `SELECT id FROM comments WHERE time_add = $1`
		if err := pg.Db.QueryRow(queryS2, CurTime).Scan(&CommentID); err != nil{
			return err
		}

		buf := model.RComment{
			CommentData: &Comment.CommentData,
			ParentID: Comment.ParentID,
			PostID: &Comment.PostID,
			NestingLevel: &tempLevel,
			TimeAdd: &TimeAdd,
			CommentID: &CommentID,
		}
		for _, val := range(occupiedSlot){
			subCh[val] <- &buf
		}
		

	}
	return nil
}

func (pg *Postgres)PostAndComment(postID *string, limit int, offset int) (*model.PostWithComment, error){
	
	query := `SELECT * FROM posts WHERE id = $1`
	var ID, Author, TimeAdd, Subject string
	var IsCommentEnbale bool
	if err := pg.Db.QueryRow(query, *postID).Scan(&ID, &Author, &TimeAdd, &IsCommentEnbale, &Subject); err != nil{
		return nil, err
	}
	var PWC = model.PostWithComment{
	Post : &model.Post{
		ID: &ID,
		Author: &Author,
		TimeAdd: &TimeAdd,
		IsCommentEnable: &IsCommentEnbale,
		Subject: &Subject,
		},
	}
	if !*PWC.Post.IsCommentEnable{
		return &PWC, nil
	}
	query = `SELECT * FROM comments WHERE post_id = $1`
	rows, err := pg.Db.Query(query, *postID)
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		var temp model.RComment
		err = rows.Scan(&temp.CommentID, &temp.CommentData, &temp.ParentID, &temp.PostID, &temp.NestingLevel, &temp.TimeAdd)
		if err != nil{
			return nil, err
		}
		PWC.Comments = append(PWC.Comments, &temp)
	}
	res := make([]*model.RComment, 0)
	for i := 0; i < len(PWC.Comments) ;i++{
		if *PWC.Comments[i].ParentID == "0"{
			res = pg.GetChild(PWC.Comments[i:], *PWC.Comments[i].CommentID, res)
		} else{
			break
		}
		
	}
	PWC.Comments = res[offset:offset + limit]
	
	return &PWC, nil
}

func(pg *Postgres) GetChild(posts []*model.RComment, id string, res []*model.RComment) []*model.RComment{
	res = append(res, posts[0])
	for i, val := range(posts){
		if *val.ParentID == id && i != 0{
			res = pg.GetChild(posts[i:], *val.CommentID, res)
		}
	}
	return res
}

func (pg *Postgres) IsCommentable(postID *string,)(bool, error){
	query := `SELECT * FROM posts WHERE id = $1`
	var ID, Author, TimeAdd,  Subject string
	var IsCommentEnbale bool
	if err := pg.Db.QueryRow(query, postID).Scan(&ID, &Author, &TimeAdd, &IsCommentEnbale, &Subject); err != nil{
		return false, err
	}
	return IsCommentEnbale, nil
}