package api

import (
	"fmt"
	"log"
	"myapi/pkg/models"
	"myapi/pkg/repository"
	"myapi/pkg/tools"
)

func TestDB() {
	config := tools.New()

	db, err := repository.New(config.Postgres.ConnStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	users, err := db.GetUsers()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range users {
		fmt.Printf(
			"id: %d, login: %s, firstName: %s, lastName: %s, age: %d, roleID: %d \n",
			item.ID,
			item.Login,
			item.FirstName,
			item.LastName,
			item.Age,
			item.RoleID,
		)
	}

	user, err := db.GetUserByID(12)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf(
		"id: %d, login: %s, firstName: %s, lastName: %s, age: %d, roleID: %d \n",
		user.ID,
		user.Login,
		user.FirstName,
		user.LastName,
		user.Age,
		user.RoleID,
	)

	err = db.NewUser(models.User{FirstName: "Артручик", Login: "GLAVNIY", LastName: "LOL", Age: 16, RoleID: 1})
	if err != nil {
		log.Fatal(err.Error())
	}
}
