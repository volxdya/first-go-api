package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Login     string
	RoleID    int
	Age       int
}

type Role struct {
	User  string
	Admin string
}
