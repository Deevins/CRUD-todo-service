package repository

import (
	"fmt"
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{
		db: db,
	}
}

func (r *TodoListPostgres) CreateList(userID int, list entity.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1, $2) RETURNING ID", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING ID", usersListsTable)

	_, err = tx.Exec(createUsersListQuery, userID, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userID int) ([]entity.TodoList, error) {
	var lists []entity.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userID)

	return lists, err
}

func (r *TodoListPostgres) GetByID(userID, listID int) (entity.TodoList, error) {
	var list entity.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userID, listID)

	return list, err
}

func (r *TodoListPostgres) Delete(userID, listID int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListsTable, usersListsTable)

	_, err := r.db.Exec(query, userID, listID)

	return err
}

func (r *TodoListPostgres) Update(userID int, listID int, input entity.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argsID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsID))
		args = append(args, *input.Title)
		argsID++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argsID))
		args = append(args, *input.Description)
		argsID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argsID, argsID+1)
	args = append(args, listID, userID)

	logrus.Debugf("updated query %s", query)
	logrus.Debugf("args %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
