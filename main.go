package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"unifizeAssignment/internal/handlers"
)

func main() {
	router := gin.Default()
	handlers.LoadHandlers(router)

	public := router.Group("/api/v1")
	public.GET("/health", Health)

	port := 8080
	err := router.Run("0.0.0.0:" + strconv.Itoa(port))
	if err != nil {
		log.Panicf("Failed to start discounts server, reason: %v", err)
		return
	}

	log.Println("Started discounts server at port: " + strconv.Itoa(port))
}

func Health(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEPlain, []byte{'0'})
}
