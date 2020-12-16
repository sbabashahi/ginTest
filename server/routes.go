package server

import (
	// "net/http"
	"gin-test/controllers"
	"github.com/gin-gonic/gin"
)
var db = make(map[string]string)

// SetupRouter funcntion
// func SetupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, CreateResponse(map[string]string{"user": user, "value": value}, "", 0, 0, 200))
// 		} else {
// 			c.JSON(http.StatusOK, CreateResponse(map[string]string{"user": user, "status": "no value"}, "", 0, 0, 200))
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }


//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
	 grp1.GET("user", controllers.GetUsers)
	 grp1.POST("user", controllers.CreateUser)
	 grp1.GET("user/:id", controllers.GetUserByID)
	 grp1.PUT("user/:id", controllers.UpdateUser)
	 grp1.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}