package route

import (
	"admin/internal/handler"
	"fmt"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouter(dbConn *gorm.DB) *gin.Engine {
	router := gin.Default()

	htmlRender := multitemplate.NewRenderer()
	htmlRender.AddFromFiles("index", "web_contents/views/index.html")
	htmlRender.AddFromFiles("login", "web_contents/views/login.html")
	router.HTMLRender = htmlRender

	router.GET("/", func(c *gin.Context) {
		fmt.Println("/에 왔음")
		c.HTML(
			http.StatusOK,
			"login",
			gin.H{
				"title": "login gogo",
			},
		)
	})
	router.GET("/user", func(c *gin.Context) {
		fmt.Println("/user에 왔음")
		c.String(200, "Hello, World!")
	})

	router.GET("/user_list", func(c *gin.Context) {
		fmt.Println("/user_list에 왔음")
		users := handler.GetTotalUserList()
		c.HTML(http.StatusOK, "./web_contents/views/templete/list_user.html", users)
	})

	router.Static("/assets", "./web_contents/assets")

	return router
}
