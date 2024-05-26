package postgres

import (
	custom_errors "comments_service/errors"
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct{
	Db *sql.DB
}

func (pg *Postgres) Register(login, hashToken string) error{
	query := `INSERT INTO users(login, token_hash) VALUES($1, $2)`
	_, err := pg.Db.Exec(query, login, hashToken)
	if err != nil{
		return custom_errors.ErrAlreadyRegistered
	}


	return nil
}


func (pg *Postgres) IsRegister(string) (bool, error){
	return true, nil
}