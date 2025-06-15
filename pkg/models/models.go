package models

type User struct {
	ID        int    `json: "id"`
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Login     string `json: "login"`
	RoleID    int    `json: "role_id"`
	Age       int    `json: "age"`
	RoleName  string `json: role_name`
}

type Role struct {
	ID       int    `json: "id"`
	RoleName string `json: "role_name"`
}
