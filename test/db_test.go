package Test

import (
	"comments_service/config"
	"comments_service/internal/storage/postgres"
	"math/rand"
	"testing"
	"time"

)

func Connect() (*postgres.Postgres, error){
	cfg := config.PostgresConfig{
		DbName : "commentsdb",
		User : "apuha",
		Password : "12345678",
		Port : "5432",
	}
	pg, err := postgres.InitDb(cfg)
	if err != nil{
		return nil, err
	}
	return pg, nil
}

func TestRegister(t *testing.T){
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano()) 
	s := make([]rune, 12)

	pg, err := Connect()
	if err != nil{
		t.Error(err)
	}
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
	err = pg.Register(string(s))
	if err != nil{
		t.Error(err)
	}
}

