package core

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"nft_platform/global"
	"nft_platform/middleware"
	"nft_platform/router"
	"nft_platform/validation"
	"os"
	"os/signal"
	"time"
)

// Run 运行Server
func Run() {
	server := newServer()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", global.Conf.Server.Port),
		Handler: server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("list:%s\n", err)
		}
	}()

	stopServer(&srv)

}

// 初始化server
func newServer() *gin.Engine {
	r := gin.New()
	gin.Default()
	bindValidator()
	setServerMode()
	setServerLogger()
	serverInfo()
	registerMiddleware(r)
	registerStatic(r)
	registerRouter(r)
	return r
}

// 停止server 回收资源
func stopServer(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("ShutDown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	global.Rdb.Close()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

// 注册路由地址
func registerRouter(e *gin.Engine) {
	router.NewRouter(e)
}

// 注册静态资源
func registerStatic(e *gin.Engine) {
	e.Static("/static", "./resource/static")
	e.Static("/public", "./resource/public")
}

// 注册中间件
func registerMiddleware(e *gin.Engine) {
	e.Use(middleware.Cors(), middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CheckSign())
	e.Use(middleware.JWT())
}

func serverInfo() {
	info := `*******************************************
*******************************************
           _.._        ,-------------------.
        ,'      \.    ( 启动成功！)
       /  __) __\ \\    \-,-----------------'
      (  (\-\(-')  ) _.-'
      /)  \\  = /  (
     /'    |--' .  \\
    (  ,---|  \-.)__\
     )(  \-.,--'   _\-.
    '/,'          (  Uu\,
     (  @       ,    \/,-' )
     \.__,  : @  /  /\--'
       |     \--'  |
       \   \-._   /
        \\        (
        /\\ .      \\. 
       / |\ \\     ,-\\
      /  \\| .)   /   \\
     ( ,'|\\    ,'     :
     | \\,\.\--\/      }
     \,'    \\  |,'    /
    / \-._   \-/      |
    \-.   \-.,'|     ;
   /        _/[\---'\\]
  :        /  |\-     '
  '           |      /
*******************************************
*******************************************`
	info += fmt.Sprintf("\n")
	info += fmt.Sprintf("server: http://127.0.0.1:%s\n", global.Conf.Server.Port)
	fmt.Println(info)
}

func setServerMode() {
	if global.Conf.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func setServerLogger() {
	if global.Conf.Logger.Type != "file" {
		return
	}
	filename := fmt.Sprintf("%s/gin.log", global.Conf.Logger.Path)
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Errorf("gin 日志创建出错，err:%s", err))
	}
	gin.DefaultWriter = io.MultiWriter(file)
}

// 自定义验证器
func bindValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", validation.Mobile)
		v.RegisterValidation("cradno", validation.CardNo)
		v.RegisterValidation("chinese", validation.Chinese)
	}
}
