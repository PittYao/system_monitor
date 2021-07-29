package main

import (
	"log"
	"os"
	"os/signal"
	"sjw_system_monitor/http"
	"syscall"
)

func main() {
	go http.ServeHTTP()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()
	log.Println("Server Start Awaiting Signal")
	<-done
	log.Println("Exiting")

}
