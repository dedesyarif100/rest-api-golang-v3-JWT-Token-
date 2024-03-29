package main

import (
	"rest-api-golang-v3/config"
	"rest-api-golang-v3/controller"
	"rest-api-golang-v3/middleware"
	"rest-api-golang-v3/repository"
	"rest-api-golang-v3/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db				*gorm.DB					= config.SetupDatabaseConnection()
	userRepository	repository.UserRepository	= repository.NewUserRepository(db)
	jwtService		service.JWTService			= service.NewJWTService()
	userService		service.UserService			= service.NewUserService(userRepository)
	authService		service.AuthService			= service.NewAuthService(userRepository)
	authController 	controller.AuthController	= controller.NewAuthController(authService, jwtService)
	userController 	controller.UserController	= controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run()
}