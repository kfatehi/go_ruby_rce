package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ruby := r.Group("/ruby")

	ruby.POST("/validate", func(ctx *gin.Context) {
		rubyCmd := exec.Command("ruby", "validate.rb")
		rubyCmd.Env = os.Environ()
		rubyCmd.Env = append(rubyCmd.Env, "SCRIPT_FILE=validate.rb")
		rubyOut, err := rubyCmd.Output()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.Data(http.StatusOK, "application/json", rubyOut)
		}
	})

	r.Run("0.0.0.0:8080")
}
