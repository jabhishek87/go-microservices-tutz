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
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /tasks [get]
func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, "List of Tasks")
}

// @Summary Find task by ID
// @Description get the given task
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /tasks/{id} [get]
func getTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "requested id : "+id)
}

// @Summary Add a new task
// @Description
// @operationId: "postTask"
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Router /tasks [post]
func postTask(c echo.Context) error {
	return c.String(http.StatusOK, "Create Task")
}

// @Summary Update a existing task
// @Description
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /tasks/{id} [put]
func putTask(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, "updated task id : "+id)
}

// @Summary Deletes a existing task
// @Description
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /tasks/{id} [delete]
func delTask(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
