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
	"strings"
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

	signature := ctx.GetHeader("X-Hub-Signature-256")
	if signature == "" {
		log.Println("signature do not have")
		ctx.Abort()
		return
	}

	go func() {
		if util.VerifySignature(signature, resByte) {
			// 先答复，要不然 github 一直发信息

			var output []byte
			path := os.Getenv(projectPath) + res.Repository.Name
			sn := os.Getenv(shellName)

			_, err = os.Stat("./git.sh")
			if os.IsExist(err) {
				os.Remove("./git.sh")
			}

			var file *os.File
			file, err = os.OpenFile("./git.sh", os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				log.Println("open git.sh failed,err:", err)
				return
			}

			file.Write([]byte(fmt.Sprintf("cd %s && git pull\n", path)))
			// 文件夹可执行
			file.Write([]byte(fmt.Sprintf("chmod +x ../%s\n", path)))
			// 文件可执行
			file.Write([]byte(fmt.Sprintf("chmod +x %s\n", sn)))
			// 执行文件
			file.Write([]byte(fmt.Sprintf("./%s\n", sn)))
			file.Close()

			// git pull && 执行脚本
			cmd := exec.Command("sh", "-c", "./git.sh")
			stdout, _ := cmd.StdoutPipe()
			stderr, _ := cmd.StderrPipe()

			err = cmd.Start()
			if err != nil {
				log.Println("starting command failed, err:", err)
				return
			}

			go asyncLog(stdout)
			go asyncLog(stderr)

			defer func() {
				err1 := recover()
				if err1 != nil {
					panic(err1)
				}
			}()

			err = cmd.Wait()
			if err != nil {
				log.Println("waiting for command execution failed, err:", err)
				return
			}

			log.Println("successful result")
			fmt.Println(string(output))

			log.Println("build successful!")

			return
		}
		log.Println("signature error!")
	}()

	ctx.JSON(200, "build successful!")
}

func asyncLog(reader io.ReadCloser) error {
	cache := ""
	buf := make([]byte, 1024, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF || strings.Contains(err.Error(), "closed") {
				err = nil
			}
			return err
		}
		if num > 0 {
			oByte := buf[:num]
			oSlice := strings.Split(string(oByte), "\n")
			line := strings.Join(oSlice[:len(oSlice)-1], "\n")
			fmt.Printf("%s%s\n", cache, line)
			cache = oSlice[len(oSlice)-1]
		}
	}
}
