package ip

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ip struct {
	Host string `json:"host"`
}

var hostIP = ip{Host: "1.2.3.4"}

func getIP(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hostIP)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi. %v !", name)
	return message
}

func StartService() {
	router := gin.Default()
	router.GET("/", getIP)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
