package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseData struct {
	Code resCode     `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, status int, code resCode, msg string, data interface{}) {
	t := &responseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	log.Printf("%+v\n", t)
	c.JSON(status, t)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, http.StatusOK, codeSuccess, "", data)
}

func ResponseError(c *gin.Context, status int, code resCode) {
	Response(c, status, code, code.Msg(), nil)
}
