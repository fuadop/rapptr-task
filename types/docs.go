package types

type ApiResponse[K interface{}] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    K      `json:"data"`
}

type DocRegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DocAddTodo struct {
	Text string `json:"text"`
}

type DocLoginRes struct {
	Token string `json:"token"`
}

type DocAddTodoRes struct {
	Created bool `json:"created"`
}

type DocEditTodoRes struct {
	Editted bool `json:"editted"`
}

type DocDeleteTodoRes struct {
	Deleted bool `json:"deleted"`
}
