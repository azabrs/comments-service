package jwt

import (
	"testing"

)

func TestCreateToken(t *testing.T) {
	type args struct {
		signingKey string
		login      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "correct JWTkey", args: args{login: "login", signingKey: "secretkey"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateToken(tt.args.signingKey, tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCheckUserToken(t *testing.T) {
	type args struct {
		tokenString string
		signingKey  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "chekUser", 
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY5MTUwOTQsIkxvZ2luIjoic2FzaGEifQ.pRDo540tPkyUSLCCh5n3JEYhzcmv0ps-2f_hvrX04KI", 
				signingKey: "secretkey",
				}, 
			want: "sasha", 
			wantErr: false,
		},

		{
			name: "NotToken", 
			args: args{
				tokenString: "asd", 
				signingKey: "secretkey",
				}, 
			wantErr: true,
		},

		{
			name: "InvalidToken", 
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc1MjIyMjQsIkxvZ2luIjoic2FzaGEifQ.xTUff0TBWRsw4E61RuP_SQ4Qo3Cae2zJcx3eiCjvuGs", 
				signingKey: "secretkey",
				}, 
			wantErr: true,
		},
		
		{
			name: "InvalidClaims", 
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTc1MjMzMjUsIk5hbWUiOiJuZXNhc2hhIn0.9RAuWyNiHaoawO_0PeUq_f7GtEGl0ymhmVNuRITFj4g", 
				signingKey: "secretkey",
				}, 
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUserToken(tt.args.tokenString, tt.args.signingKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUserToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckUserToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
