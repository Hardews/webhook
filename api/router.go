/**
 * @Author: Hardews
 * @Date: 2023/10/8 23:48
 * @Description:
**/

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const redirectUrl = "https://hardews.cn"

func Init() {
	r := gin.Default()

	r.POST("/shell", shellBuild)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
	})

	r.Run(":8090")
}
