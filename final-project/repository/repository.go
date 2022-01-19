package repository

import (
	"final-project/db"
	"final-project/entity"
	"final-project/utils/error_formats"
	"final-project/utils/error_utils"
)

type RepositoryInterface interface {
	Create(*entity.TodoList) (*entity.TodoList, error_utils.MessageErr)
	Update(*entity.TodoList) (*entity.TodoList, error_utils.MessageErr)
	Delete(int)
	GetAll() (*[]entity.TodoList, error_utils.MessageErr)
	GetById(int) (*entity.TodoList, error_utils.MessageErr)
}

type repository struct{}

var Repo RepositoryInterface = &repository{}

func (r *repository) Create(todoReq *entity.TodoList) (*entity.TodoList, error_utils.MessageErr) {
	db := db.NewDB()

	sql := "INSERT INTO todolist (title) values ($1) RETURNING id, title, completed, createdat"
	row := db.QueryRow(sql, todoReq.Title)

	var todo entity.TodoList

	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &todo, nil
}

func (r *repository) Update(todoReq *entity.TodoList) (*entity.TodoList, error_utils.MessageErr) {
	db := db.NewDB()

	sql := "UPDATE todolist SET title = $2, completed = $3 WHERE id = $1 RETURNING id, title, completed, createdat"
	row := db.QueryRow(sql, todoReq.Id, todoReq.Title, todoReq.Completed)

	var todo entity.TodoList

	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &todo, nil
}

func (r *repository) Delete(id int) {
	db := db.NewDB()

	sql := "DELETE FROM todolist WHERE id=$1"
	db.Exec(sql, id)
}

func (r *repository) GetAll() (*[]entity.TodoList, error_utils.MessageErr) {
	db := db.NewDB()

	sql := `SELECT id, title, completed, createdat FROM todolist`
	rows, err := db.Query(sql)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	todos := []entity.TodoList{}
	for rows.Next() {
		todo := entity.TodoList{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, error_formats.ParseError(err)
		}

		todos = append(todos, todo)

	}

	return &todos, nil
}

func (r *repository) GetById(id int) (*entity.TodoList, error_utils.MessageErr) {
	db := db.NewDB()

	sql := `SELECT id, title, completed, createdat FROM todolist WHERE id = $1`
	row := db.QueryRow(sql, id)

	var todo entity.TodoList
	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &todo, nil
}
