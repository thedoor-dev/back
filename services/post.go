package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thedoor-dev/back/db"
	"github.com/thedoor-dev/back/models"
)

var postLen int = 2

func PostLen(c *gin.Context) {
	ResponseSuccess(c, postLen)
}

func PostList(c *gin.Context) {
	data := struct {
		Page int `json:"page"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("out", err)
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	if data.Page <= 0 {
		log.Println("out")
		ResponseError(c, http.StatusUnsupportedMediaType, codePayloadError)
		return
	}
	var ps []models.PostList
	err = db.PostList(&ps, (data.Page-1)*postLen, postLen)
	if err != nil {
		log.Println("out", err)
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	for i, v := range ps {
		err = db.PostListTags(&(ps[i].Tags), v.ID)
		if err != nil {
			log.Println("out", err)
			ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
			return
		}
	}
	ResponseSuccess(c, ps)
}

func PostOne(c *gin.Context) {
	pid, err := strconv.ParseInt(c.Param("id"), 16, 64)
	if err != nil {
		ResponseError(c, http.StatusNotFound, codeParamError)
		return
	}
	var p models.Post
	isAdmin := c.GetBool("isAdmin")
	err = db.PostOne(&p, pid, isAdmin)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	ResponseSuccess(c, p)
	log.Println(pid)
}

func PostNew(c *gin.Context) {
	// var p models.Post
	// err := c.BindJSON(&p)
	// if err != nil {
	// 	log.Println("out", err)
	// 	ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
	// 	return
	// }
	// log.Printf("%+v\n", p)
	// err = db.PostNew(&p)
	// if err != nil {
	// 	log.Println("out", err)
	// 	ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
	// 	return
	// }
	// ResponseSuccess(c, nil)
	data := struct {
		Title    string   `json:"title"`
		Abstract string   `json:"abstract"`
		Article  string   `json:"article"`
		Public   bool     `json:"public"`
		Tag      []string `json:"tag"`
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		log.Println("out", err)
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	log.Printf("%+v\n", data)
	err = db.PostNewWithTag(data.Title, data.Abstract, data.Article, data.Public, data.Tag)
	if err != nil {
		log.Println("out", err)
		ResponseError(c, http.StatusInternalServerError, codeServiceBusy)
		return
	}
	ResponseSuccess(c, nil)
}
