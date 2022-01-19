package controller

import (
	"final-project/entity"
	"final-project/service"
	"final-project/utils/error_utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Tags todos
// @Description create todolist
// @ID create todolist
// @Accept json
// @Produce json
// @Param RequestBody body doc_datas.CreateTodolistRequest true "request body json"
// @Success 200 {object} doc_datas.CreateTodolistResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 422 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todos/ [post]
func TodoCreate(c *gin.Context) {
	var todos entity.TodoList

	if err := c.ShouldBindJSON(&todos); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	data, err := service.Serv.Create(&todos)

	if err != nil {
		c.JSON(err.Status(), "error at service create")
		return
	}

	todoList := entity.TodoList{
		Id:        data.Id,
		Title:     data.Title,
		Completed: data.Completed,
		CreatedAt: data.CreatedAt,
	}

	c.JSON(http.StatusOK, todoList)
}

// Update Todo godoc
// @Tags todos
// @Description update todolist
// @ID update todolist
// @Accept json
// @Produce json
// @Param RequestBody body doc_datas.UpdateTodolistRequest true "request body json"
// @Success 200 {object} doc_datas.UpdateTodolistResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 422 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todos/{id} [put]
func TodoUpdate(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		log.Fatal(err)
	}

	var todos entity.TodoList
	if err := c.ShouldBindJSON(&todos); err != nil {
		theErr := error_utils.NewBadRequest("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}

	todos.Id = id
	res, err := service.Serv.Update(&todos)

	if err != nil {
		c.JSON(0, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// Delete Todo godoc
// @Tags todos
// @Description delete todolist
// @ID delete todolist
// @Accept json
// @Produce json
// @Param id path int true "todolist id"
// @Success 200 {object} doc_datas.DeleteTodolistResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 422 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todos/{id} [delete]
func TodoDelete(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		log.Fatal(err)
	}

	service.Serv.Delete(id)

	c.JSON(http.StatusOK, "delete sukses")
}

// GetAll Todo godoc
// @Tags todos
// @Description getall todolist
// @ID getall todolist
// @Accept json
// @Produce json
// @Success 200 {array} doc_datas.GetAlldTodolistResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 422 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todos/ [get]
func TodoGetAll(c *gin.Context) {
	data, err := service.Serv.GetAll()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	// todoList := entity.TodoList{
	// 	Id:        data.Id,
	// 	Title:     data.Title,
	// 	Completed: data.Completed,
	// 	CreatedAt: data.CreatedAt,
	// }

	c.JSON(http.StatusOK, data)
}

// GetById Todo godoc
// @Tags todos
// @Description getbyid todolist
// @ID getbyid todolist
// @Accept json
// @Produce json
// @Param id path int true "todolist id"
// @Success 200 {object} doc_datas.GetByIdTodolistResponse
// @Failure 400 {object} error_utils.MessageErrData
// @Failure 422 {object} error_utils.MessageErrData
// @Failure 500 {object} error_utils.MessageErrData
// @Router /todos/{id} [get]
func TodoGetById(c *gin.Context) {
	// var todosReq entity.TodoList

	// paramId, err := todosReq.GetOrdersIdParams(c)

	paramId := c.Param("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		log.Fatal(err)
	}

	res, err := service.Serv.GetById(id)
	if err != nil {
		c.JSON(0, err)
		return
	}

	todos := entity.TodoList{
		Id:        res.Id,
		Title:     res.Title,
		Completed: res.Completed,
		CreatedAt: res.CreatedAt,
	}

	c.JSON(http.StatusOK, todos)
}
