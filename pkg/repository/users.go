package repository

import (
	"context"
	"myapi/pkg/models"
)

func (repo *PGRepo) GetUsers() ([]models.User, error) {
	rows, err := repo.pool.Query(context.Background(), "SELECT id, first_name, last_name, age, login, role_id FROM users;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var item models.User
		err = rows.Scan(
			&item.ID,
			&item.FirstName,
			&item.LastName,
			&item.Age,
			&item.Login,
			&item.RoleID,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, item)
	}

	return users, nil
}

func (repo *PGRepo) GetUserByID(id int) (models.User, error) {
	var user models.User

	err := repo.pool.QueryRow(context.Background(), "SELECT id, first_name, last_name, age, login, role_id FROM users where id = $1;", id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Login,
		&user.RoleID,
	)

	if err != nil {
		return user, err
	}

	return user, err
}

func (repo *PGRepo) NewUser(item models.User) error {
	_, err := repo.pool.Exec(
		context.Background(),
		"INSERT INTO users (first_name, last_name, age, login, role_id) VALUES ($1, $2, $3, $4, $5)",
		item.FirstName, item.LastName, item.Age, item.Login, item.RoleID,
	)

	return err
}
