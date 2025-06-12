package main

import (
	"context"
	"fmt"
	"log"
	"myapi/pkg/models"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	connStr := "postgres://postgres:qwerty@localhost:5432/damn_test"

	//api := api.New("localhost:3001", http.NewServeMux())
	//api.FillEndpoints()
	//log.Fatal(api.ListenAndServe())
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var u models.User

	err = conn.QueryRow(context.Background(), "SELECT id, first_name, last_name, login, age, role_id FROM users WHERE id = 12;").Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Login,
		&u.Age,
		&u.RoleID,
	)

	fmt.Println(u)

	if err != nil {
		log.Fatal(err.Error())
	}
}
