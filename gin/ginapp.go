package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateApp() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := r.Run("0.0.0.0:8000")
	if err != nil {
		return
	}
}
