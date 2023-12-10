package handler

import (
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
	"github.com/Jadepypy/distributed-social-media/application/internal/usecase"
	"github.com/gin-contrib/sessions"
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
	userGroup := server.Group("/user")
	userGroup.POST("/signup", u.SignUp)
	userGroup.POST("/login", u.Login)
	userGroup.POST("/edit", u.Edit)
	userGroup.GET("/profile", u.Profile)
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
	type LogInReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LogInReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	err := u.useCase.LogIn(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	sess := sessions.Default(ctx)
	sess.Set("user_id", req.Email)
	sess.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		Birthday string `json:"birthday"`
		Intro    string `json:"intro"`
		Nickname string `json:"nickname"`
	}

	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	email := ctx.Value("user_id").(string)
	err := u.useCase.Edit(ctx, domain.User{
		Email:    email,
		Birthday: req.Birthday,
		Intro:    req.Intro,
		Nickname: req.Nickname,
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

func (u *UserHandler) Profile(ctx *gin.Context) {
	email := ctx.Value("user_id").(string)
	user, err := u.useCase.GetProfile(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"birthday": user.Birthday,
		"intro":    user.Intro,
		"nickname": user.Nickname,
	})
}
