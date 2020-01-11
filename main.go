package main

import (
	
	"net/http"
	
	"fmt"
	"github.com/gin-gonic/gin"
)
// SimpleApp - Simpleapp
type SimpleApp struct{
    Appname string `json:"appname" binding:"required"`
    Apptype string `json:"apptype" binding:"required"`
}
var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.POST("/simpleapp", func(ctx *gin.Context) {
		
		//var app SimpleApp
        //ctx.BindJSON(&app)
		ctx.Request.ParseForm()
		apptype := ctx.PostForm("apptype")
		appname :=  ctx.PostForm("appname")
		applib :=  ctx.PostForm("lib")
		depmgmt := ctx.PostForm("dependency-management")

		fmt.Println(depmgmt)
		CreateFiles(appname,apptype,applib)		 
		
		ctx.Header("Content-Description", "File Transfer")
	//	ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename=app.zip" )
		ctx.Header("Content-Type", "application/zip")
		ctx.File("C:\\Users\\rajesh.deshpande02\\AppData\\Local\\Temp\\app.zip")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}



