package ip

import (
	"fmt"
	"github.com/gin-gonic/gin"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	"github.com/tavsec/gin-healthcheck/config"
	"net/http"
	"os"
	"time"
)

type ip struct {
	Host string `json:"host"`
}

var PORT = os.Getenv("PORT")

var hostIp ip

func getIP(c *gin.Context) {
	if header := c.Request.Header.Values("X-Forwarded-For")[0]; header != "" {
		hostIp = ip{Host: header}
	} else {
		hostIp = ip{Host: c.Request.RemoteAddr}
	}
	c.IndentedJSON(http.StatusOK, hostIp)
}

func StartService() {
	router := gin.Default()
	router.GET("/", getIP)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.Header,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	healthcheck.New(router, config.DefaultConfig(), []checks.Check{})
	err := router.Run("0.0.0.0:" + PORT)
	if err != nil {
		return
	}
}
