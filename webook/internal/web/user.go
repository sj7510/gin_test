package web

import (
	"gin_test/webook/internal/domain"
	"gin_test/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

// UserHandler Define user related routes
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:         svc,
		emailExp:    regexp.MustCompile(`\A([\w+\-].?)+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`, regexp.None),
		passwordExp: regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, regexp.None),
	}
}

func (u *UserHandler) RegisterUserRoutes(server *gin.Engine) {

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

	ok, err := u.emailExp.MatchString(req.Email)
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

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "password's length must be big 8 and contain letter and digit")
		return
	}

	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "email is used")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}

	ctx.String(http.StatusOK, "sign up %v\n", req)
}

// Login user login
func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "email or password error")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}

	// Jwt
	claims := UserClaims{
		Uid: user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenStr, err := token.SignedString([]byte("secret"))

	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}
	ctx.Header("x-jwt-token", tokenStr)
	ctx.String(200, "login success")
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{MaxAge: -1})
	err := sess.Save()
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}
	ctx.String(200, "logout success")
}

// Edit user info
func (u *UserHandler) Edit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "this is edit function"})
}

// Profile get user profile
func (u *UserHandler) Profile(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")

	if !exists {
		ctx.String(http.StatusOK, "system error")
		return
	}

	claimsValue, ok := claims.(*UserClaims)
	if !ok {
		ctx.String(http.StatusOK, "system error")
		return
	}
	ctx.JSON(200, gin.H{"uid": claimsValue.Uid})
}

type UserClaims struct {
	jwt.StandardClaims
	Uid int64
}
