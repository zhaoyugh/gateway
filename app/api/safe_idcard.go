package api
import (
    "code.qschou.com/gateway/app/aes"
    "github.com/gin-gonic/gin"
    "net/http"
)
const encrypt_key = "3wjyxqDPNyrd4QrhxTycRMU4dFN2lCm4"
//身份证加密
func SafeEncryptIdcard(c *gin.Context) {
    idcard := c.Query("idcard")
    payload,err := aes.OpenSslAesEncrypt(idcard, encrypt_key)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.String(http.StatusOK, "加密的字符串为: %s", payload)
    return
}

//身份证解密
func SafeDecryptIdcard(c *gin.Context){
    payload := c.Query("payload")

    idcart,err := aes.OpenSslAesDecrypt(payload, encrypt_key)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.String(http.StatusOK, "身份证号为::: %s", idcart)
    return
}