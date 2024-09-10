package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler Define user related routes
type UserHandler struct {
}

func (u *UserHandler) registerUserRoutes(server *gin.Engine) {

	ug := server.Group("/users")

	ug.POST("/signup", u.SignUp)

	ug.POST("/login", u.Login)

	ug.POST("/edit", u.Edit)

	ug.GET("/profile", u.Profile)
}

// SignUp user sign up
func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	// Regular expression check email and password
	const (
		emailRegexPattern    = `\A([\w+\-].?)+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	ok, err := emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "email regexp error")
		return
	}

	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "twice input password not equal")
		return
	}

	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	ok, err = passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "password's length must be big 8 and contain letter and digit")
		return
	}

	ctx.String(http.StatusOK, "sign up %v\n", req)
}

// Login user login
func (u *UserHandler) Login(ctx *gin.Context) {
	ctx.HTML(200, "user_login.html", nil)
}

// Edit user info
func (u *UserHandler) Edit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "this is edit function"})
}

// Profile get user profile
func (u *UserHandler) Profile(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "this is profile function"})
}
