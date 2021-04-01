package main

import "github.com/gin-gonic/gin"

func CpuIntensiveJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cpu!",
	})
}

func IoIntensiveJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "io!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/cpuintensivejob", CpuIntensiveJob)
	r.GET("/iointensivejob", IoIntensiveJob)
	r.Run(":3000")
}
