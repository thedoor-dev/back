package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var imgNameID *SnowFlake

func init() {
	imgNameID = &SnowFlake{}
}

func SaveImgToGitee(img []byte) string {
	imgName := fmt.Sprintf("%x", imgNameID.GetVal())

	urlPost := "https://gitee.com/api/v5/repos/SunspotsInys/demo/contents/thedoor/" + imgName + ".png"
	data := make(map[string]string)
	data["access_token"] = "dce1a797e5ba27c727fa092155642cb6"
	data["content"] = base64.StdEncoding.EncodeToString(img)
	data["message"] = "upload img " + imgName
	data["branch"] = "master"
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("out")
		return ""
	}
	resp, err := http.Post(urlPost, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("out")
		return ""
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("out")
		return ""
	}
	log.Printf("%s\n", respBody)
	return imgName + ".png"
}
