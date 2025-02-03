package main

import (
	"golang_starter_template/pkg/server"
	"golang_starter_template/pkg/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Initialize the logger
	utils.Init()
	utils.LoggerInfo.Println("Starting server...")

	// Signals Channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// handle signals
	go func() {
		<-signals
		utils.CleanUp()
		os.Exit(0)
	}()

	// Rotate logs
	utils.Rotate()

	// Start the server
	err := server.Start(os.Args[1:])
	if err != nil {
		log.Default().Println(err)
		utils.LoggerError.Println(utils.Error+err.Error(), utils.Reset)
		return
	}
	select {}
}
