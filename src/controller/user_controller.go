package controller

import (
	"net/http"
	"sample/src/data/request"
	"sample/src/data/response"
	"sample/src/helper"
	"sample/src/service"

	"github.com/gin-gonic/gin"
    "github.com/satori/go.uuid"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Save(c *gin.Context) {
	CreateUserRequest := request.CreateUserRequest{}
	err := c.ShouldBindJSON(&CreateUserRequest)
	helper.CheckErr(err)

	controller.UserService.Save(CreateUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Update(c *gin.Context) {
	UpdateUserRequest := request.UpdateUserRequest{}
	err := c.ShouldBindJSON(&UpdateUserRequest)
	helper.CheckErr(err)

	userId := c.Param("id")
	id, err := uuid.FromString(userId)
	helper.CheckErr(err)
	UpdateUserRequest.Id = id

	controller.UserService.Update(UpdateUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) Delete(c *gin.Context) {
	userId := c.Param("id")
	id, err := uuid.FromString(userId)
	helper.CheckErr(err)
	
	controller.UserService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}	

func (controller *UserController) FindById(c *gin.Context) {
	userId := c.Param("id")
	id, err := uuid.FromString(userId)
	helper.CheckErr(err)

	result := controller.UserService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) FindAll(c *gin.Context) {
	result := controller.UserService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
