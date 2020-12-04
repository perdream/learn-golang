package main

import (
	"fmt"
	"net/http"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func main() {
	var router = routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//开始监听服务，监听TCP网络地址，Addr和调用应用程序处理连接上的请求
	/*
	   func (srv *Server) ListenAndServe() error {
	       addr := srv.Addr
	       if addr == "" {
	           addr = ":http"
	       }
	       ln, err := net.Listen("tcp", addr)
	       if err != nil {
	           return err
	       }
	       return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
	   }
	*/
	s.ListenAndServe()
}
