package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func AES256CBCEncrypt(plantText string, key string, iv []byte) (value string, err error) {
	plaintext := PKCS7Padding([]byte(plantText))
	ciphertext := make([]byte, len(plaintext))
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func AES256CBCDecrypt(plantText string, key string, iv string) (value string, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return
	}
	var iiv []byte
	if iiv, err = base64.StdEncoding.DecodeString(iv); err != nil {
		return
	}

	var cipherText []byte
	if cipherText, err = base64.StdEncoding.DecodeString(plantText); err != nil {
		return
	}

	mode := cipher.NewCBCDecrypter(block, iiv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText = PKCS7UnPadding(cipherText)
	return string(cipherText), nil
}

func getHmacCode(s string, key []byte) string {
	h := hmac.New(sha256.New, key)
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type Crypt struct {
	Iv    string `json:"iv"`
	Value string `json:"value"`
	Mac   string `json:"mac"`
}

func OpenSslAesEncrypt(value string, key string) (payload string, err error) {
	iv := make([]byte, 16)
	_, err = rand.Read(iv)

	var encode string
	if encode, err = AES256CBCEncrypt(value, key, iv); err != nil {
		return
	}

	ivv := base64.StdEncoding.EncodeToString(iv)

	crypt := Crypt{Iv: ivv, Value: encode, Mac: getHmacCode(ivv+encode, []byte(key))}

	var bs []byte
	if bs, err = json.Marshal(crypt); err != nil {
		return
	}

	return base64.StdEncoding.EncodeToString(bs), nil
}

func OpenSslAesDecrypt(payload string, key string) (value string, err error) {
	var bs []byte
	if bs, err = base64.StdEncoding.DecodeString(payload); err != nil {
		return
	}
	crypt := Crypt{}
	if err = json.Unmarshal(bs, &crypt); err != nil {
		return
	}

	return AES256CBCDecrypt(crypt.Value, key, crypt.Iv)
}
