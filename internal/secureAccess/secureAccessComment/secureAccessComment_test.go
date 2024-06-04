package secure_access_comment

import (
	"comments_service/graph/model"
	"comments_service/internal/storage/mocks"
	storage "comments_service/internal/storage/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSecureAccessComment_Authorize(t *testing.T) {
	type fields struct {
		JWTKey string
	}
	type args struct {
		login string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "correct JWTkey", 
			fields: fields{
				JWTKey: "secretkey",
				}, 
			args: args{
				login: "login",
				}, 
			wantErr: false,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := SecureAccessComment{
				JWTKey: tt.fields.JWTKey,
			}
			_, err := auth.Authorize(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureAccessComment.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestSecureAccessComment_Authentication(t *testing.T) {
	type fields struct {
		JWTKey string
	}
	type args struct {
		IdentificationData model.IdentificationData
		stor               storage.Storage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid",
			fields: fields{
					JWTKey: "secretkey",
			},
			args: args{
				IdentificationData: model.IdentificationData{
					Login: "sasha",
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4ODgwMDUsIkxvZ2luIjoic2FzaGEifQ.dgUOGwLbSQPHujNSotwEIChJ79E0bWtRSFkW_w58niU",
				},
				stor: *mocks.NewStorage(t),
			},
			wantErr: false,
		},

		{
			name: "Invalid_login_in_token",
			fields: fields{
					JWTKey: "secretkey",
			},
			args: args{
				IdentificationData: model.IdentificationData{
					Login: "petr",
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTY4ODgwMDUsIkxvZ2luIjoic2FzaGEifQ.dgUOGwLbSQPHujNSotwEIChJ79E0bWtRSFkW_w58niU",
				},
				stor: *mocks.NewStorage(t),
			},
			wantErr: true,
		},

		{
			name: "User_already_registred",
			fields: fields{
					JWTKey: "secretkey",
			},
			args: args{
				IdentificationData: model.IdentificationData{
					Login: "petr",
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUzMTc1MzQ3ODAsIkxvZ2luIjoicGV0ciJ9.Q_JzcVD62rK3rFIPASdXtjadnRHrozsTUNNccTv7oVw",
				},
				stor: *mocks.NewStorage(t),
			},
			wantErr: true,
		},

		{
			name: "User_already_registred",
			fields: fields{
					JWTKey: "secretkey",
			},
			args: args{
				IdentificationData: model.IdentificationData{
					Login: "petr",
					Token: "adsdas",
				},
				stor: *mocks.NewStorage(t),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.args.stor.On("IsLoginExist", "sasha").Return(nil)
		tt.args.stor.On("IsLoginExist", mock.AnythingOfType("string")).Return(fmt.Errorf("error"))
		t.Run(tt.name, func(t *testing.T) {
			auth := SecureAccessComment{
				JWTKey: tt.fields.JWTKey,
			}
			if err := auth.Authentication(tt.args.IdentificationData, &tt.args.stor); (err != nil) != tt.wantErr {
				t.Errorf("SecureAccessComment.Authentication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
