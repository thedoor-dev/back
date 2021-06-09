package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thedoor-dev/back/configs"
	"github.com/thedoor-dev/back/db"
	"github.com/thedoor-dev/back/models"
	"github.com/thedoor-dev/back/routes"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	// 读取配置文件路径
	if len(os.Args) < 2 {
		log.Println("需要指定配置文件eg: ./thedoor config.yaml")
		return
	}

	// 初始化配置
	err := configs.Init(os.Args[1])
	if err != nil {
		log.Printf("配置文件初始化失败，程序异常退出%v ", err)
		return
	}
	// 初始化数据库 - MariaDB连接
	err = db.Init(configs.Conf.MariaDBConfig)
	if err != nil {
		log.Printf("MariaDB初始化连接失败，程序异常退出%v", err)
		return
	}
	if configs.Conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	e := routes.Init()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", configs.Conf.Port),
		Handler: e,
	}

	// 优雅关机
	go func() {
		// err := srv.ListenAndServeTLS("/www/nginx/ssl/aovj.cert.pem", "/www/nginx/ssl/aovj.key.pem")
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("关闭服务中...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Printf("服务关闭失败: %v\n", err)
		return
	}
	log.Println("服务已成功退出")
}
func init() {
	log.SetFlags(log.Llongfile)
	var ts models.TagArr
	ts = append(ts, models.Tag{
		ID:   0,
		PID:  0,
		Name: "avc",
	})
	ts = append(ts, models.Tag{
		ID:   0,
		PID:  0,
		Name: "abv",
	})
	a, _ := json.Marshal(ts)
	log.Printf("%s\n", a)
}
