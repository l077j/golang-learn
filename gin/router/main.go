package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.GET("/findUser/:username/:userid", FindUser)
	e.GET("/downloadFile/*filepath", UserPage)

	log.Fatalln(e.Run(":8080"))
}

// 命名参数示例
// c.Param() 方法获取路由中的两个参数username 和 userid。这些参数会被分配到 username 和 userid 变量中，并使用 c.String() 方法将它们作为响应返回客户端
func FindUser(c *gin.Context) {
	username := c.Param("username")
	userid := c.Param("userid")
	c.String(http.StatusOK, "username is %s\n userid is %s", username, userid)
}

// 路径参数示例
func UserPage(c *gin.Context) {
	filepath := c.Param("filepath")
	c.String(http.StatusOK, "filepath is %s", filepath)
}
