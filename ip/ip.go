package ip

import (
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
	"net/http"
	"os"
)

type ip struct {
	Host string `json:"host"`
}

var PORT = os.Getenv("PORT")

var hostIp ip

func getIP(c *gin.Context) {
	if header := c.Request.Header.Get("X-Forwarded-For"); header != "" {
		hostIp = ip{Host: header}
	} else {
		hostIp = ip{Host: c.Request.RemoteAddr}
	}
	c.IndentedJSON(http.StatusOK, hostIp)
}

func StartService() {
	router := gin.Default()
	router.GET("/", getIP)

	healthcheck.New(router, config.DefaultConfig(), []checks.Check{})

	err := router.Run("0.0.0.0:" + PORT)
	if err != nil {
		return
	}
}
