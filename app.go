package main

import (
	"context"
	"fmt"
	"os"
	"github.com/vijaya2911/server"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	//go abortServer(cancel)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "5555"
	}
	chatSer := chatserver.NewChatServer("0.0.0.0", port, ctx)
	errc := chatSer.Start()
	err:=<-errc
	fmt.Printf("endGame: chatServer error=%s\n",err)
}


func abortServer(cancel context.CancelFunc) {
	os.Stdin.Read(make([]byte, 1))
	cancel()
}
