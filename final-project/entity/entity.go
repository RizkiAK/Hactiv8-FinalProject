package entity

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoList struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (todos *TodoList) GetOrdersIdParams(c *gin.Context) (int, error) {
	paramId := c.Param("ordersId")

	// fmt.Println("ini Id orders =>", paramId)

	ordersId, err := strconv.Atoi(paramId)

	if err != nil {
		return 0, err
	}

	return ordersId, nil
}
