package Routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello, world!")
}

func Router() {
	router := gin.Default()
	router.GET("/test", test)

	router.POST("/post", func(c *gin.Context) {

		type CamposPost struct {
			ID      int    `json:"id"`
			Page    int    `json:"page"`
			Name    string `json:"name"`
			Message string `json:"message"`
		}

		var postData CamposPost

		if err := c.ShouldBindJSON(&postData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("id: %d; page: %d; name: %s; message: %s\n", postData.ID, postData.Page, postData.Name, postData.Message)

		c.JSON(200, gin.H{
			"status": "success",
			"data":   postData,
		})
	})
	router.Run("localhost:8080")
}
