package postgres

import (
	"comments_service/config"
	custom_errors "comments_service/errors"
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


func (pg *Postgres) IsRegister(string) (bool, error){
	return true, nil
}