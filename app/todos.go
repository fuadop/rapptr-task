package app

import (
	"database/sql"
	"net/http"

	"github.com/fuadop/rapptr-task/types"
	"github.com/gin-gonic/gin"
)

// @Summary		Get all todos
// @Description	Get all todos of a user
// @Tags			todos
// @Produce		json
// @Success		200	{object}	types.ApiResponse[[]types.Todo]
// @Failure		400	{object}	types.ApiResponse[any]
// @Failure		500	{object}	types.ApiResponse[any]
// @Router			/todos [get]
// @Security		BearerAuth
func (a *App) GetTodos(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		return
	}

	todos, err := a.Model.GetTodos(id.(string))
	if err != nil && err != sql.ErrNoRows {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	a.HandleResponseJSON(c, http.StatusOK, "Todos fetched", todos, nil)
}

// @Summary		Get a todo
// @Description	Get a todo
// @Tags			todos
// @Produce		json
// @Param			id	path		string	true	"todo id"
// @Success		200	{object}	types.ApiResponse[types.Todo]
// @Failure		400	{object}	types.ApiResponse[any]
// @Failure		404	{object}	types.ApiResponse[any]
// @Failure		500	{object}	types.ApiResponse[any]
// @Router			/todos/{id} [get]
// @Security		BearerAuth
func (a *App) GetTodo(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}

	id := c.Param("id")
	todo, err := a.Model.GetTodo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
			return
		}

		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	if todo.UserID != userId.(string) { // don't disclose other users todos to others
		a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
		return
	}

	a.HandleResponseJSON(c, http.StatusOK, "Todo fetched", todo, nil)
}

// @Summary		Add a todo
// @Description	Add a todo
// @Tags			todos
// @Accept			json
// @Produce		json
// @Param			payload	body		types.DocAddTodo	true	"Request body"
// @Success		200		{object}	types.ApiResponse[types.DocAddTodoRes]
// @Failure		400		{object}	types.ApiResponse[any]
// @Failure		500		{object}	types.ApiResponse[any]
// @Router			/todos [post]
// @Security		BearerAuth
func (a *App) AddTodo(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}

	var todo types.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		a.HandleResponseJSON(c, http.StatusBadRequest, "Validation Error", nil, err)
		return
	}

	if err := a.Model.InsertTodo(todo.Text, userId.(string)); err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Request failed", nil, err)
		return
	}

	res := map[string]bool{"created": true}
	a.HandleResponseJSON(c, http.StatusCreated, "Todo created", res, nil)
}

// @Summary		Change a todo
// @Description	Change a todo
// @Tags			todos
// @Produce		json
// @Param			id		path		string				true	"todo id"
// @Param			payload	body		types.DocAddTodo	true	"Request body"
// @Success		200		{object}	types.ApiResponse[types.DocEditTodoRes]
// @Failure		400		{object}	types.ApiResponse[any]
// @Failure		404		{object}	types.ApiResponse[any]
// @Failure		500		{object}	types.ApiResponse[any]
// @Router			/todos/{id} [patch]
// @Security		BearerAuth
func (a *App) EditTodo(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}

	id := c.Param("id")

	var todo types.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		a.HandleResponseJSON(c, http.StatusBadRequest, "Validation Error", nil, err)
		return
	}

	t, err := a.Model.GetTodo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
			return
		}

		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	if t.UserID != userId.(string) { // don't disclose other users todos to others
		a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
		return
	}

	if err = a.Model.UpdateTodo(id, todo.Text); err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Request failed", nil, err)
		return
	}

	res := map[string]bool{"updated": true}
	a.HandleResponseJSON(c, http.StatusOK, "Todo updated", res, nil)
}

// @Summary		Delete a todo
// @Description	Delete a todo.
// @Tags			todos
// @Produce		json
// @Param			id	path		string	true	"id"
// @Success		200	{object}	types.ApiResponse[types.DocDeleteTodoRes]
// @Failure		404	{object}	types.ApiResponse[any]
// @Failure		500	{object}	types.ApiResponse[any]
// @Router			/todos/{id} [delete]
// @Security		BearerAuth
func (a *App) DeleteTodo(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		return
	}

	id := c.Param("id")

	t, err := a.Model.GetTodo(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
			return
		}

		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	if t.UserID != userId.(string) { // don't disclose other users todos to others
		a.HandleResponseJSON(c, http.StatusNotFound, "Todo not found", nil, nil)
		return
	}

	if err := a.Model.DeleteTodo(id); err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	res := map[string]bool{"deleted": true}
	a.HandleResponseJSON(c, http.StatusOK, "Todo deleted", res, nil)
}
