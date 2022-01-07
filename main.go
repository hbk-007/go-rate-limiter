package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	/*Simple gin router for demo purpose*/
	router := gin.Default()
	InitRoutes(router)

	err := router.Run(":" + "8080")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
