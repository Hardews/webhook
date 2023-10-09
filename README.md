# webhook
处理 github webhook 发送的网络请求，然后执行相关脚本。



## 环境变量

使用本程序需要设置的环境变量有：

- **PROJECT_BASE_PATH**，项目的基础路径。
- **WEBHOOK_SECRET**，你设置的 webhook secret。
- **PROJECT_SHELL_NAME**，收到请求后，需要执行 shell 脚本的名称。



## 使用

**本机需要拥有 go 环境**

```
git clone https://github.com/Hardews/webhook.git
```

将仓库 clone 下来后

```go
go mod tidy
```

然后

```
go run main.go
```

