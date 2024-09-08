package web

import "github.com/gin-gonic/gin"

// UserHandler Define user related routes
type UserHandler struct {
}

// SignUp user sign up
func (u *UserHandler) SignUp(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "this is sign up function"})
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
