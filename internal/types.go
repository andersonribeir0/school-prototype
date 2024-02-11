package internal

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role []Role `json:"role"`
}

type Role string

const (
	Student Role = "Student"
	Teacher Role = "Teacher"
	Admin   Role = "Admin"
)
