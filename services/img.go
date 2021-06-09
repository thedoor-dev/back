package services

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedoor-dev/back/utils"
)

func ImgUpload(c *gin.Context) {
	fh, err := c.FormFile("file[]")
	if err != nil {
		log.Println("out")
		ResponseError(c, http.StatusForbidden, codePayloadError)
		return
	}
	file, err := fh.Open()
	if err != nil {
		log.Println("out")
		ResponseError(c, http.StatusForbidden, codePayloadError)
		return
	}
	defer file.Close()
	img, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("out")
		ResponseError(c, http.StatusForbidden, codePayloadError)
		return
	}
	imgName := utils.SaveImgToGitee(img)
	if imgName == "" {
		ResponseError(c, http.StatusBadRequest, codeServiceBusy)
		return
	}
	log.Println("上传成功")
	ResponseSuccess(c, imgName)
}
