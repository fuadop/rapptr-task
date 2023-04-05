package app

import (
	"net/http"

	"github.com/fuadop/rapptr-task/types"
	"github.com/gin-gonic/gin"
)

// @Summary		Login a user
// @Description	Login a user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			payload	body		types.DocRegisterUser	true	"Request body"
// @Success		200		{object}	types.ApiResponse[types.DocLoginRes]
// @Failure		400		{object}	types.ApiResponse[any]
// @Failure		500		{object}	types.ApiResponse[any]
// @Router			/auth/login [post]
func (a *App) Login(c *gin.Context) {
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		a.HandleResponseJSON(c, http.StatusBadRequest, "Validation Error", nil, err)
		return
	}

	u, err := a.Model.GetUserByUsername(user.Username)
	if err != nil {
		a.HandleResponseJSON(c, http.StatusUnauthorized, "Incorrect username or password", nil, err)
		return
	}

	hash := a.Model.HashPassword(user.Password) // hash input password & compare
	if hash != u.Password {
		a.HandleResponseJSON(c, http.StatusUnauthorized, "Incorrect username or password", nil, nil)
		return
	}

	token, err := a.GenerateJWTAccessToken(u.ID, u.Username)
	if err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal Server Error", nil, err)
		return
	}

	res := map[string]string{"token": token}
	a.HandleResponseJSON(c, http.StatusOK, "Login successful", res, nil)
}

// @Summary		Register a user
// @Description	Register a user
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			payload	body		types.DocRegisterUser	true	"Request body"
// @Success		200		{object}	types.ApiResponse[any]
// @Failure		400		{object}	types.ApiResponse[any]
// @Failure		500		{object}	types.ApiResponse[any]
// @Router			/auth/register [post]
func (a *App) Register(c *gin.Context) {
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		a.HandleResponseJSON(c, http.StatusBadRequest, "Validation Error", nil, err)
		return
	}

	if err := a.Model.InsertUser(user.Username, user.Password); err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Registeration failed", nil, err)
		return
	}

	a.HandleResponseJSON(c, http.StatusCreated, "Registeration complete, please login", nil, nil)
}

// @Summary		Get the current logged in user
// @Description	Get the current logged in user
// @Tags			users
// @Produce		json
// @Success		200	{object}	types.ApiResponse[types.User]
// @Failure		400	{object}	types.ApiResponse[any]
// @Failure		500	{object}	types.ApiResponse[any]
// @Router			/auth/me [get]
// @Security		BearerAuth
func (a *App) GetMe(c *gin.Context) {
	id, ok := c.Get("userId")
	if !ok {
		// a.HandleResponseJSON(c, http.StatusUnauthorized, "Unauthorized", nil, errors.New("unauthorized"))
		return
	}

	u, err := a.Model.GetUser(id.(string))
	if err != nil {
		a.HandleResponseJSON(c, http.StatusInternalServerError, "Internal server error", nil, err)
		return
	}

	u.Password = "" // hide password from response
	a.HandleResponseJSON(c, http.StatusOK, "Fetched loggedin user", u, nil)

}
