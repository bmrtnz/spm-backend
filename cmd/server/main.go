package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialise le routeur Gin avec les paramètres par défaut
	router := gin.Default()

	// Définition de la route /healthcheck
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "SPM Backend is running!",
		})
	})

	// Démarrage du serveur sur le port 8080 par défaut
	// Vous pouvez changer le port si nécessaire, par exemple : router.Run(":9000")
	err := router.Run()
	if err != nil {
		// Idéalement, utiliser un logger structuré ici (sera fait dans UST0.12)
		panic("Failed to start Gin server: " + err.Error())
	}
}