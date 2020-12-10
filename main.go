package main

import (
    "fmt"
    "github.com/fvbock/endless"
    "log"
    "mayGo/models"
    "mayGo/pkg/logging"
    "mayGo/routers"
    "syscall"
    
    "mayGo/pkg/setting"
)

func main() {
    setting.Setup()
    models.Setup()
    logging.Setup()
    
    endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
    
    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }
    
    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}
