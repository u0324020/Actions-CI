package main

import (
	"github.com/u0324020/ACTIONS-CI/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context){
	var(
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != ""{
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("should input username")})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != ""{
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("should input PW")})
		return
	}
	if err := auth.Auth(username, password); err == nil{
		c.HTML(http.StatusOK, "login.html", gin.H{"success": "login!"})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}

}

func main(){
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":8888")

}
