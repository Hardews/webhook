/**
 * @Author: Hardews
 * @Date: 2023/10/8 23:52
 * @Description:
**/

package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"os/exec"
	"webhook/model"
	"webhook/util"
)

var (
	projectPath = "PROJECT_BASE_PATH"  // 项目基础路径，比如 /www/
	shellName   = "PROJECT_SHELL_NAME" // shell 脚本的名称，如 build.sh
)

func shellBuild(ctx *gin.Context) {
	resByte, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	var res model.JsonBody
	json.Unmarshal(resByte, &res)

	// 一些鉴权，看个人怎么鉴权咯
	// 这里使用 github 给的 secret

	log.Println("webhook result")
	fmt.Println(res)

	signature := ctx.GetHeader("X-Hub-Signature")
	if signature == "" {
		log.Println("signature do not have")
		ctx.Abort()
		return
	}

	if util.VerifySignature(signature, resByte) {
		var output []byte
		path := os.Getenv(projectPath) + res.Repository.Name
		sn := os.Getenv(shellName)

		// git pull
		cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s && git pull", path))
		if output, err = cmd.Output(); err != nil {
			log.Println("git pull error, err:", err)
			ctx.Abort()
			return
		}

		log.Println("git pull result")
		fmt.Println(string(output))

		// 开始执行脚本
		cmd = exec.Command("sh", "-c", fmt.Sprintf("cd %s && ./%s", path, sn))
		if output, err = cmd.Output(); err != nil {
			log.Println("build error, err:", err)
			ctx.Abort()
			return
		}

		log.Println("build result")
		fmt.Println(string(output))

		log.Println("build successful!")
		return
	}
	log.Println("signature error!")
}
