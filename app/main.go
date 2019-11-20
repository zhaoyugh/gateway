package main

import "fmt"
import "github.com/gin-gonic/gin"

func main(){
    fmt.Println("赵宇大好人11222")
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong1",
        })
    })
    r.Run(":8081")
    fmt.Println("赵宇大好人111")
}

