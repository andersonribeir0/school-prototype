package models

import (
	"encoding/hex"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		err     error
	}{
		{
			name: "Valid user",
			user: User{
				Username: "testuser",
				Password: "password",
				Roles:    []Role{Student},
			},
			wantErr: false,
		},
		{
			name: "No username",
			user: User{
				Password: "password",
				Roles:    []Role{Student},
			},
			wantErr: true,
			err:     ErrUsernameRequired,
		},
		{
			name: "No password",
			user: User{
				Username: "testuser",
				Roles:    []Role{Student},
			},
			wantErr: true,
			err:     ErrPwdRequired,
		},
		{
			name: "No roles",
			user: User{
				Username: "testuser",
				Password: "password",
			},
			wantErr: true,
			err:     ErrAtLeastOneRoleRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != tt.err {
				t.Errorf("Validate() error = %v, expected %v", err, tt.err)
			}
		})
	}
}

func TestUser_HasRole(t *testing.T) {
	user := User{
		Username: "testuser",
		Password: "password",
		Roles:    []Role{Student, Teacher},
	}

	tests := []struct {
		name string
		role Role
		want bool
	}{
		{
			name: "Has role",
			role: Student,
			want: true,
		},
		{
			name: "Does not have role",
			role: Admin,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := user.HasRole(tt.role); got != tt.want {
				t.Errorf("HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ParseAndCheckPwd(t *testing.T) {
	user := User{
		Username: "testuser",
		Password: "password",
	}

	if err := user.ParsePwd(); err != nil {
		t.Fatalf("ParsePwd() failed: %v", err)
	}

	// Verify password is hashed
	_, err := hex.DecodeString(user.Password)
	if err != nil {
		t.Errorf("Password not properly hashed: %v", err)
	}

	// Check password
	if !user.CheckPwd("password") {
		t.Errorf("CheckPwd() failed to validate the correct password")
	}

	// Check wrong password
	if user.CheckPwd("wrongpassword") {
		t.Errorf("CheckPwd() incorrectly validated a wrong password")
	}
}
