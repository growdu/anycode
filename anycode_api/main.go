package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

// DatabaseConfig stores the database configuration
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

// AppConfig stores the application-wide configurations
type AppConfig struct {
	Database DatabaseConfig `json:"database"`
}

// User represents the user model
type User struct {
	tableName struct{} `pg:"users"`
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
}

// Global variable to hold the DB connection
var db *pg.DB

func main() {
	args := os.Args
	var url string
	if len(args) > 1 {
		url = args[1]
	} else {
		fmt.Println("unkown input, using ", args[0], "ip:port.")
		return
	}
	// Load config
	configFile, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var config AppConfig
	if err = json.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// Initialize PostgreSQL connection
	db = pg.Connect(&pg.Options{
		Addr:     config.Database.Host + ":" + config.Database.Port,
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.DBName,
	})
	defer db.Close()

	// 创建一个默认的 Gin 引擎
	r := gin.Default()
	// 允许所有来源访问，并且允许使用 POST 方法
	r.Use(cors.Default())

	// Login route
	r.POST("/login", loginHandler)

	r.POST("/change", updatePasswordHandler) // 新增修改密码的路由

	r.POST("/action", actionHandler)
	// 启动 HTTP 服务器，监听在 8080 端口
	r.Run(url)
}

func loginHandler(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(200, gin.H{"error": "Invalid JSON provided", "ok": 0})
		return
	}

	// Check user in the database
	user := &User{}
	err := db.Model(user).Where("username = ?", u.Username).Select()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "User not found", "ok": 0})
		return
	}

	if user.Password != u.Password {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect password", "ok": 0})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "", "ok": 1})
}

// 新增修改密码的处理函数
func updatePasswordHandler(c *gin.Context) {
	var u struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid JSON provided", "ok": 0})
		return
	}

	user := &User{}
	err := db.Model(user).Where("username = ?", u.Username).Select()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "User not found", "ok": 0})
		return
	}

	// 验证旧密码是否正确
	if user.Password != u.OldPassword {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect old password", "ok": 0})
		return
	}

	// 更新密码
	user.Password = u.NewPassword
	_, err = db.Model(user).Column("password").WherePK().Update()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to update password", "ok": 0})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "", "ok": 1})
}

// 新增 action 处理函数
func actionHandler(c *gin.Context) {
	var req struct {
		Action string `json:"action"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid JSON provided", "ok": 0})
		return
	}

	var cmd *exec.Cmd
	switch req.Action {
	case "start", "stop", "restart", "upgrade", "program":
		// 执行命令，将动作传递给二进制程序
		cmd = exec.Command("./node.sh", req.Action)
	default:
		c.JSON(400, gin.H{"error": "Invalid action provided"})
		return
	}

	// 执行命令
	if err := cmd.Run(); err != nil {
		c.JSON(500, gin.H{"error": "Failed to execute action"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "", "ok": 1})
}
