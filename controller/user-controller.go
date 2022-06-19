package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"rest-api-golang-v3/dto"
	"github.com/gin-gonic/gin"
	"rest-api-golang-v3/helper"
	"rest-api-golang-v3/service"
	"github.com/golang-jwt/jwt/v4"
)

// USER CONTROLLER IS A
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService	service.UserService
	jwtService	service.JWTService
}

// NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService: jwtService,
	}
}

func (user *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("FAILED TO PROCESS REQUEST", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := user.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := user.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (user *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := user.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	userProfile := user.userService.Profile(id)
	res := helper.BuildResponse(true, "OK!", userProfile)
	context.JSON(http.StatusOK, res)
}
