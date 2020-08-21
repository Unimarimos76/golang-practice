package controller

import (
	"net/http"
	"time"
	"webapp/src/model"

	"github.com/gin-gonic/gin"
)

func TasksGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER by id desc1")
	if err != nil {
		panic(err.Error())
	}

	tasks := []model.Task{}
	for result.Next() {
		task := model.Task{}
		var id uint
		var createdAt, updatedAt time.Time
		var title string

		if err = result.Scan(&id, &createdAt, &updatedAt, &title); err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
