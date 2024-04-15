package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "fmt"
)

var count uint64


func main() {
    args := os.Args
    var url string
    if len(args) > 1 {
        url = args[1]
    } else {
        fmt.Println("unkown input, using ", args[0], "ip:port.");
        return
    }
    // 创建一个默认的 Gin 引擎
    r := gin.Default()
    r.Use(Cors())
    // 定义一个路由，处理 GET 请求
    r.GET("/login", func(c *gin.Context) {
        // 返回一个 JSON 响应
        count++
        c.JSON(http.StatusOK, gin.H{
            "visit":200,
        })
    })

    // 启动 HTTP 服务器，监听在 8080 端口
    r.Run(url)
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method
        origin := c.Request.Header.Get("Origin") //请求头部
        if origin != "" {
            //接收客户端发送的origin （重要！）
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            //服务器支持的所有跨域请求的方法
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
            //允许跨域设置可以返回其他子段，可以自定义字段
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
            // 允许浏览器（客户端）可以解析的头部 （重要）
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
            //设置缓存时间
            c.Header("Access-Control-Max-Age", "172800")
            //允许客户端传递校验信息比如 cookie (重要)
            c.Header("Access-Control-Allow-Credentials", "true")
        }

        //允许类型校验
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "ok!")
        }

        defer func() {
            if err := recover(); err != nil {
                println("Panic info is: %v", err)
            }
        }()

        c.Next()
    }
}


