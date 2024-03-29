package repository

import (
	"fmt"
	entity "github.com/deevins/todo-restAPI/internal/entities"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listID uuid.UUID, item entity.TodoItem) (uuid.UUID, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return Nil, err
	}
	//var itemID int
	itemID := uuid.NewV4()
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING ID", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)

	err = row.Scan(&itemID)
	if err != nil {
		tx.Rollback()
		return Nil, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listID, itemID)
	if err != nil {
		tx.Rollback()
		return Nil, err
	}
	return itemID, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userID, listID uuid.UUID) ([]entity.TodoItem, error) {
	var items []entity.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
													  INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetItemByID(userID, itemID uuid.UUID) (entity.TodoItem, error) {
	var item entity.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
													  INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	//query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
	//								INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
	//	todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemID, userID); err != nil {
		return item, err
	}
	return item, nil
}
