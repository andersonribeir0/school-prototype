package models

import (
	"testing"
)

func TestUser_HasRole(t *testing.T) {
	user := User{
		Username: "john_doe",
		Roles:    []Role{Student, Admin},
	}

	tests := []struct {
		role Role
		want bool
	}{
		{Student, true},
		{Teacher, false},
		{Admin, true},
	}

	for _, tt := range tests {
		t.Run(string(tt.role), func(t *testing.T) {
			if got := user.HasRole(tt.role); got != tt.want {
				t.Errorf("User.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr error
	}{
		{"ValidUser", User{Username: "john_doe", Password: "secret", Roles: []Role{Student}}, nil},
		{"NoUsername", User{Password: "secret", Roles: []Role{Student}}, ErrUsernameRequired},
		{"NoPassword", User{Username: "john_doe", Roles: []Role{Student}}, ErrPwdRequired},
		{"NoRoles", User{Username: "john_doe", Password: "secret"}, ErrAtLeastOneRoleRequired},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.user.Validate(); err != tt.wantErr {
				t.Errorf("User.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_ParsePwd(t *testing.T) {
	user := User{
		Username: "john_doe",
		Password: "secret",
	}

	want := "2bb80d5..."
	if err := user.ParsePwd(); err != nil {
		t.Errorf("User.ParsePwd() error = %v", err)
	}

	if user.Password[:10] != want {
		t.Errorf("User.ParsePwd() got = %v, want %v", user.Password[:10], want)
	}
}
