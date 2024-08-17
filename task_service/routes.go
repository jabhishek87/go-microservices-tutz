package main

import (
	"cmp"
	"net/http"

	_ "task/docs" // Import the docs

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// getRoot godoc
//
//	@Summary		Show the status of server.
//	@Description	get the status of server.
//	@Tags			root
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response
//	@Router			/ [get]
func getRoot(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		response{Code: http.StatusOK, Message: "Server is up and running"},
	)

}

// @Summary		Get Tasks
// @Description	get All tasks
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]Task
// @Router			/api/v1/tasks [get]
func getTasks(c echo.Context) error {
	var tasks []Task
	db.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

// @Summary		Find task by uuid
// @Description	get the given task
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			uuid	path		int	true	"uuid"
// @Success		200		{object}	string
// @Router			/api/v1/tasks/{uuid} [get]
func getTask(c echo.Context) error {
	uuid := c.Param("uuid")
	var task Task

	res := db.First(&task, "uuid = ?", uuid) // find product with uuid
	if res.Error != nil {
		return c.JSON(
			http.StatusNotFound,
			response{Code: http.StatusNotFound, Message: "record not found"},
		)
	}
	return c.JSON(http.StatusOK, task)
	// return c.JSON(http.StatusOK, "requested id : "+id)
}

// @Summary	Add a new task
// @Description
// @Tags			Tasks
// @operationId:	"postTask"
// @Accept			json
// @Produce		json
// @Param body body Task true "body of the request"
// @Success		200	{object}	string
// @Router			/api/v1/tasks [post]
func postTask(c echo.Context) error {
	t := new(Task)
	if err := c.Bind(t); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	newTask := Task{
		UUID:   uuid.New(),
		Title:  t.Title,
		Status: cmp.Or(t.Status, "new"),
		UserID: cmp.Or(t.UserID, 0),
	} // Todo{Title: title, Status: status, UUID: uuid.New()}
	db.Create(&newTask)
	return c.JSON(http.StatusCreated, newTask)
}

// @Summary	Update a existing task
// @Description
// @Tags		Tasks
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"id"
// @Success	200	{object}	string
// @Router		/api/v1/tasks/{id} [put]
func putTask(c echo.Context) error {
	uuid := c.Param("uuid")

	var task, bodyTask Task
	// search for records if exists
	if res := db.First(&task, "uuid = ?", uuid); res.Error != nil {
		return c.JSON(
			http.StatusNotFound,
			response{Code: http.StatusNotFound, Message: "record not found"},
		)
	}
	// validate request body first
	// https://echo.labstack.com/docs/binding
	if err := c.Bind(&bodyTask); err != nil {
		return err
	}
	db.Model(&task).Updates(bodyTask)
	return c.JSON(http.StatusAccepted, task)
}

// @Summary	Deletes a existing task
// @Description
// @Tags		Tasks
// @Produce	json
// @Param		id	path		int	true	"id"
// @Success	200	{object}	string
// @Router		/api/v1/tasks/{id} [delete]
func delTask(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
