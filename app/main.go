package main

import "fmt"
import "github.com/gin-gonic/gin"

func main(){
    fmt.Println("赵宇大好人1122ee3332")
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "赵宇大好人",
        })
    })
    r.Run(":8081")
    fmt.Println("赵宇大好人111")
}

