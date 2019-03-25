package main

import (
	"context"
	"fmt"
	"os"
	"github.com/vijaya2911/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go abortServer(cancel)
	chatSer := chatserver.NewChatServer("localhost", "5555", ctx)
	errc := chatSer.Start()
	err:=<-errc
	fmt.Printf("endGame: chatServer error=%s\n",err)
}


func abortServer(cancel context.CancelFunc) {
	os.Stdin.Read(make([]byte, 1))
	cancel()
}
