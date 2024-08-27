package controller

import (
	"GoSimpleAPI/src/modules/todo/dto"
	"GoSimpleAPI/src/modules/todo/service"
	"GoSimpleAPI/src/modules/utils/response/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoController struct {
	todoService *service.TodoService
}

func NewTodoController(todoService *service.TodoService) *TodoController {
	return &TodoController{todoService}
}

// CreateTodo
// @Summary      Create a new Todo item
// @Description  This endpoint allows the user to create a new Todo item.
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        todo  body      dto.CreateTodoDTO  true  "Todo Data"
// @Success      201   {object}  entity.Todo        "Successfully created Todo item"
// @Failure      400   {object}  error.Response     "Invalid input"
// @Router       /todo [post]
func (c *TodoController) CreateTodo(ctx *gin.Context) {
	var createTodoDTO dto.CreateTodoDTO
	if err := ctx.ShouldBindJSON(&createTodoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: err.Error()})
		return
	}

	todo := c.todoService.CreateTodo(createTodoDTO)
	ctx.JSON(http.StatusCreated, todo)
}

// UpdateTodo
// @Summary      Update an existing Todo item
// @Description  This endpoint allows the user to update an existing Todo item by ID.
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        id      path      int     true  "Todo ID"  // The ID of the Todo item to update
// @Param        todo    body      dto.UpdateTodoDTO  true  "Updated Todo Data"
// @Success      200     {object}  entity.Todo         "Successfully updated Todo item"
// @Failure      400     {object}  error.Response      "Invalid input"
// @Failure      404     {object}  error.Response      "Todo not found"
// @Failure      500     {object}  error.Response      "Internal server error"
// @Router       /todo/{id} [put]
func (c *TodoController) UpdateTodo(ctx *gin.Context) {
	var updateTodoDTO dto.UpdateTodoDTO
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: "Invalid ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&updateTodoDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: err.Error()})
		return
	}

	todo, err := c.todoService.UpdateTodo(id, updateTodoDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Response{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

// GetToDoList
// @Summary      Get a list of Todo items
// @Description  This endpoint retrieves a list of Todo items, optionally filtered by a search query.
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        search  query     string  false  "Search query to filter Todo items"
// @Success      200     {array}   entity.Todo      "Successfully retrieved list of Todo items"
// @Failure      500     {object}  error.Response    "Internal server error"
// @Router       /todo [get]
func (c *TodoController) GetToDoList(ctx *gin.Context) {
	search := ctx.Query("search")

	todos := c.todoService.GetTodoList(search)
	ctx.JSON(http.StatusOK, todos)
}

// GetToDoById
// @Summary      Get a Todo item by ID
// @Description  This endpoint retrieves a Todo item based on the provided ID.
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        id      path      int     true  "Todo ID"  // The ID of the Todo item to retrieve
// @Success      200     {object}  entity.Todo         "Successfully retrieved Todo item"
// @Failure      400     {object}  error.Response      "Invalid ID"
// @Failure      404     {object}  error.Response      "Todo not found"
// @Failure      500     {object}  error.Response      "Internal server error"
// @Router       /todo/{id} [get]
func (c *TodoController) GetToDoById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: "Invalid ID"})
		return
	}

	todo := c.todoService.GetTodoById(id)
	ctx.JSON(http.StatusOK, todo)
}
