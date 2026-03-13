package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	jsonV2()
}

func serverTest() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	serverStart()
	<-sigChan
	fmt.Println("shutdown")
}
