package Test

import (
	inmemory "comments_service/internal/storage/inMemory"
	"testing"
	"math/rand"
	"time"
)

func TestRegisterMemory(t *testing.T){
	
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano()) 
	s := make([]rune, 12)
	for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
	im := inmemory.InitMemory()
	if err := im.Register(string(s)); err != nil{
		t.Error(err)
	}
}
