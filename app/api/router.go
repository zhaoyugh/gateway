package api

import "github.com/gin-gonic/gin"

func Router(){
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "赵宇大好人",
        })
    })
    r.GET("/safe_encrypt_idcard",SafeEncryptIdcard)
    r.GET("/safe_decrypt_idcard",SafeDecryptIdcard)
    r.Run(":9099")
}
