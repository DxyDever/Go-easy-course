package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟私有数据
var secrets = gin.H{
	"ls":   gin.H{"email": "ls@lianshiclass.com", "phone": "123456"},
	"yang": gin.H{"email": "yang@lianshiclass.com", "phone": "111111"},
	"edu":  gin.H{"email": "edu@lianshiclass.com", "phone": "666666"},
}

func main() {
	r := gin.New()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"ls":   "123",
		"yang": "111",
		"edu":  "666",
		"lucy": "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取提交的⽤户名（AuthUserKey）
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET "})
		}
	})
	r.Run(":8080")
}
