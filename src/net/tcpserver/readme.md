### TCP服务器 ###

模仿gevent的StreamServer封装golang的tcp server 


sample:

    package main
    
    import (
    	"fmt"
    	"net"
    	"net/tcpserver"
    )
    
    type Bot struct {
    	tcpserver.EndPoint
    	RecvBox chan []byte
    }
    
    func (bot *Bot) OnData(data []byte) {
    	fmt.Println("Recv:", string(data))
    	bot.RecvBox <- data
    }
    
    func (bot *Bot) OnConnectionLost(err error) {
    	fmt.Println("Connection Lost:", err.Error())
    	bot.Ctrl <- false
    }
    
    func (bot *Bot) Handle() {
    	for {
    		select {
    		case data := <-bot.RecvBox:
    			bot.PutData(data)
    			// to do something
    		}
    	}
    }
    
    func connectionHandler(conn *net.TCPConn) {
    	recvBox := make(chan []byte, 12)
    	bot := &Bot{RecvBox: recvBox}
    	bot.Init(conn, 10, 16, bot.OnData, bot.OnConnectionLost)
    
    	go bot.Handle()
    
    	bot.Start()
    }
    
    func main() {
    	server := tcpserver.NewStreamServer(":7005", connectionHandler)
    	server.Start()
    }