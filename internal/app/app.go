package app

import (
	"log"
	"os"
	"os/signal"
	"qsoft/internal/handler"
	"qsoft/internal/server"
	"qsoft/internal/service"
	"syscall"
)

func Run() {
	s := service.NewYearService()
	h := handler.NewYearHandler(s)
	srv := server.NewServer(h)

	go func() {
		err := srv.Run()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
