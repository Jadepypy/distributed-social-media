package handler

import (
	"fmt"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	fmt.Println("receive!!!2")

	userGroup := server.Group("/user")
	userGroup.POST("/signup", u.SignUp)
	userGroup.POST("/login", u.Login)
	userGroup.POST("/edit", u.Edit)
	userGroup.POST("/profile", u.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirm_password"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// TODO: validate request

	if strings.Compare(req.Password, req.ConfirmPassword) != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "password does not match",
		})
		return
	}

	err := u.useCase.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (u *UserHandler) Login(ctx *gin.Context) {
}

func (u *UserHandler) Edit(ctx *gin.Context) {
}

func (u *UserHandler) Profile(ctx *gin.Context) {
}
