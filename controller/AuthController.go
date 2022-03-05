package controller

import (
	"github.com/3langn/learn-go/dto"
	"github.com/3langn/learn-go/helper"
	"github.com/3langn/learn-go/models"
	"github.com/3langn/learn-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// https://stackoverflow.com/questions/25382073/defining-golang-struct-function-using-pointer-or-not
type IAuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthController struct {
}
var(
	authService = new(service.AuthService)
	jwtService = new(service.JwtService)
)
// @BasePath /api

// @Summary Login
// @Schemes
// @Description User login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginDTO true "Login"
// @Success 200
// @Router /auth/login [post]
func (a AuthController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := c.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObject{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// @Summary Register
// @Schemes
// @Description User resgister
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterDTO true "Register"
// @Success 200
// @Router /auth/register [post]
func (a AuthController) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := c.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObject{})
		c.JSON(http.StatusConflict, response)
	} else {
		createdUser := authService.CreateUser(registerDTO)
		token := jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		c.JSON(http.StatusCreated, response)
	}
}
