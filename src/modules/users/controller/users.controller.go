package controller

import (
	"GoSimpleAPI/src/modules/users/dto"
	"GoSimpleAPI/src/modules/users/service"
	"GoSimpleAPI/src/modules/utils/response/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

// CreateUser
// @Summary      Create a new User item
// @Description  This endpoint allows the user to create a new User item.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserDTO  true  "User Data"
// @Success      201   {object}  entity.User        "Successfully created User item"
// @Failure      400   {object}  error.Response     "Invalid input"
// @Router       /user [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	var createUserDTO dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: err.Error()})
		return
	}

	user := c.userService.CreateUser(createUserDTO)
	ctx.JSON(http.StatusCreated, user)
}

// UpdateUser
// @Summary      Update an existing User item
// @Description  This endpoint allows the user to update an existing User item by ID.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id      path      int     true  "User ID"  // The ID of the User item to update
// @Param        user    body      dto.UpdateUserDTO  true  "Updated User Data"
// @Success      200     {object}  entity.User         "Successfully updated User item"
// @Failure      400     {object}  error.Response      "Invalid input"
// @Failure      404     {object}  error.Response      "User not found"
// @Failure      500     {object}  error.Response      "Internal server error"
// @Router       /user/{id} [put]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var updateUserDTO dto.UpdateUserDTO
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: "Invalid ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&updateUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: err.Error()})
		return
	}

	user, err := c.userService.UpdateUser(id, updateUserDTO)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.Response{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetUserList
// @Summary      Get a list of User items
// @Description  This endpoint retrieves a list of User items, optionally filtered by a search query.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        search  query     string  false  "Search query to filter User items"
// @Success      200     {array}   entity.User      "Successfully retrieved list of User items"
// @Failure      500     {object}  error.Response    "Internal server error"
// @Router       /user [get]
func (c *UserController) GetUserList(ctx *gin.Context) {
	search := ctx.Query("search")

	users := c.userService.GetUserList(search)
	ctx.JSON(http.StatusOK, users)
}

// GetUserById
// @Summary      Get a User item by ID
// @Description  This endpoint retrieves a User item based on the provided ID.
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id      path      int     true  "User ID"  // The ID of the User item to retrieve
// @Success      200     {object}  entity.User         "Successfully retrieved User item"
// @Failure      400     {object}  error.Response      "Invalid ID"
// @Failure      404     {object}  error.Response      "User not found"
// @Failure      500     {object}  error.Response      "Internal server error"
// @Router       /user/{id} [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.Response{Error: "Invalid ID"})
		return
	}

	user := c.userService.GetUserById(id)
	ctx.JSON(http.StatusOK, user)
}
