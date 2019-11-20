package main

import "fmt"
import "github.com/gin-gonic/gin"

func main(){
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run()
    fmt.Println("赵宇大好人111")
}

