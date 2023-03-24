package repository

import (
	"fmt"
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (uuid.UUID, error) {
	//var id int
	id := uuid.NewV4()
	query := fmt.Sprintf("INSERT INTO %s (id, name, username, password_hash) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, id, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return Nil, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
