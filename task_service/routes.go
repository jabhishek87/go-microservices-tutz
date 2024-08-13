package main

import (
	"net/http"

	_ "task/docs" // Import the docs

	"github.com/labstack/echo/v4"
)

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Task Service! up and running")
}

// @Summary Get Tasks
// @Description get All tasks
// @Produce  json
// @Success 200 {object} string
// @Router /tasks [get]
func getTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "requested id : "+id)
}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "List of Tasks")
}

func postTask(c echo.Context) error {
	return c.String(http.StatusOK, "Create Task")
}

func putTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "updated task id : "+id)
}

func delTask(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
