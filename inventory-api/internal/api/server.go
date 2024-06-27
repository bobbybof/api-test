package api

import (
	"github.com/bobbybof/inventory-api/config"
	"github.com/bobbybof/inventory-api/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  repository.Store
	Router *gin.Engine
}

func NewServer(config config.Config, store repository.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.getRoutes()
	return server, nil
}

func (server *Server) getRoutes() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "X-XSRF-TOKEN"},
		AllowMethods: []string{"POST", "PUT", "GET", "DELETE", "PATCH"},
	}))

	router.GET("user", server.GetUserByEmail)
	router.POST("user", server.CreateUser)

	router.POST("product", server.CreateProduct)
	router.GET("products", server.GetAllProducts)

	server.Router = router
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
