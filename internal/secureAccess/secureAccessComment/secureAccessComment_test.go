package secure_access_comment

import (
	"comments_service/graph/model"
	"comments_service/internal/storage"
	"testing"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := SecureAccessComment{
				JWTKey: tt.fields.JWTKey,
			}
			got, err := auth.Authorize(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureAccessComment.Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SecureAccessComment.Authorize() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := SecureAccessComment{
				JWTKey: tt.fields.JWTKey,
			}
			if err := auth.Authentication(tt.args.IdentificationData, tt.args.stor); (err != nil) != tt.wantErr {
				t.Errorf("SecureAccessComment.Authentication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
