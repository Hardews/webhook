/**
 * @Author: Hardews
 * @Date: 2023/10/8 23:48
 * @Description:
**/

package api

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()

	r.POST("/shell", shellBuild)

	r.Run(":8090")
}
