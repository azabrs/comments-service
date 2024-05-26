package postgres

import (
	"comments_service/config"
	custom_errors "comments_service/errors"
	"comments_service/internal/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct{
	Db *sql.DB
}

func InitDb(c config.PostgresConfig) (Postgres, error){
	str := fmt.Sprintf("dbname=%s user=%s password=%s port=%s sslmode = disable", c.DbName, c.User, c.Password, c.Port)
	db, err := sql.Open("postgres", str)
	if err != nil{
		return Postgres{},  fmt.Errorf("cant connect to db, %v", err)
	}
	return Postgres{Db: db}, nil
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
