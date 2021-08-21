package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuhalang/stupidbot/internal/app/service"
)

// Server serves HTTP requests for our bot service.
type Server struct {
	eventService *service.EventService
	router       *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(eventService *service.EventService) (*Server, error) {

	server := &Server{
		eventService: eventService,
	}

	server.setupRouter()
	return server, nil
}

// func NewServerSSL(eventService *service.EventService) {

// 	x509, errTls := tls.LoadX509KeyPair(os.Getenv("SSLCRT"), os.Getenv("SSLKEY"))
// 	if errTls != nil {
// 		fmt.Println(errTls)
// 	}
// 	fmt.Println(os.Getenv("SSLCRT"), "SSLKEY")
// 	var server *http.Server
// 	server = &http.Server{
// 		Addr:    eventService.Config.ServerAddress,
// 		Handler: r,
// 		TLSConfig: &tls.Config{
// 			Certificates: []tls.Certificate{x509},
// 		},
// 	}
// 	server.ListenAndServeTLS("", "")
// }

func (server *Server) setupRouter() {
	router := gin.Default()

	// Ping handler
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/webhook", server.handleEvent)
	router.GET("/webhook", server.verifyToken)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
