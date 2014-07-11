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
    }
    
    func (bot *Bot) OnData(data []byte) {
    	fmt.Println("Recv:", string(data))
    	bot.PutData(data)
    }
    
    func (bot *Bot) OnConnectionLost(err error) {
    	fmt.Println("Connection Lost:", err.Error())
    }
    
    func connectionHandler(conn *net.TCPConn) {
    	bot := &Bot{}
    	bot.Init(conn, 10, 16, bot.OnData, bot.OnConnectionLost)
    	bot.Start()
    }
    
    func main() {
    	server := tcpserver.NewStreamServer(":7005", connectionHandler)
    	server.Start()
    }