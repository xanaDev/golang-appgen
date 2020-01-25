package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//WebServer struct
type WebServer struct {
	// gin server reference
	server *gin.Engine
	// channel for cleanup notification
	cleanUp chan int
}

//Create this will load default server config
func Create(cleanup chan int) *WebServer {
	server := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	server.Use(cors.Default())
	server.Use(gin.Recovery())

	ginServer := &WebServer{
		server:  server,
		cleanUp: cleanup,
	}
	return ginServer
}

func (ws *WebServer) Run(port string) {
	ws.server.Run(port)
}
