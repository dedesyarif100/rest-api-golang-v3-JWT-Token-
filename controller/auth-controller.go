package controller

import (
	"net/http"
	"rest-api-golang-v3/dto"
	"rest-api-golang-v3/entity"
	"rest-api-golang-v3/helper"
	"rest-api-golang-v3/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AUTH CONTROLLER INTERFACE IS A CONTRACT WHAT THIS CONTROLLER CAN DO
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authControllerStruct struct {
	authService	service.AuthService
	jwtService	service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authControllerStruct{
		authService: authService,
		jwtService: jwtService,
	}
}

func (auth *authControllerStruct) Login(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "HELLO LOGIN",
	// })

	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("FAILED TO PROCESS REQUEST", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := auth.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := auth.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("PLEASE CHECK AGAIN YOUR CREDENTIAL", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (auth *authControllerStruct) Register(ctx *gin.Context) {
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"message": "HELLO REGISTER",
	// })

	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("FAILED TO PROCESS REQUEST", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !auth.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("FAILED TO PROCESS REQUEST", "DUPLICATE EMAIL", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := auth.authService.CreateUser(registerDTO)
		token := auth.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}