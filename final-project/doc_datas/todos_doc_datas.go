package doc_datas

import "time"

type CreateTodolistRequest struct {
	Title string `json:"title" example:"Makan"`
}

type CreateTodolistResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Makan"`
	Completed bool      `json:"completed" example:"false"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-18T11:45:59.136128+07:00"`
}

type UpdateTodolistRequest struct {
	Title     string `json:"title" example:"Makan (Sudah)"`
	Completed bool   `json:"completed" example:"true"`
}

type UpdateTodolistResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Makan (Sudah)"`
	Completed bool      `json:"completed" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-18T11:45:59.136128+07:00"`
}

type GetByIdTodolistResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Makan (Sudah)"`
	Completed bool      `json:"completed" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-18T11:45:59.136128+07:00"`
}

type GetAlldTodolistResponse struct {
	Id        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Makan (Sudah)"`
	Completed bool      `json:"completed" example:"true"`
	CreatedAt time.Time `json:"created_at" example:"2022-01-18T11:45:59.136128+07:00"`
}

type DeleteTodolistResponse struct {
	Message string `json:"message" example:"delete sukses"`
}
