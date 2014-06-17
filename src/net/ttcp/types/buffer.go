package types

// 处理好的,要发送的数据,调用Send, 把数据添加到pending队列,然后由一个协程专门负责发送

import (
	"fmt"
	"net"
)

type SenderBuffer struct {
	ctrl    chan bool   // receive exit signal
	pending chan []byte // pending Packet
	max     int32       // max queue size
	conn    net.Conn    // connection
	sess    *Session    // session
}

// packet sender goroutine
func (buf *SenderBuffer) Start() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("caught panic in buffer goroutine")
			panic(x)
		}
	}()

	for {
		select {
		case data := <-buf.pending:
			buf.rawSend(data)
		case <-buf.ctrl: // session end, send final data, !重要:如果buf.ctrl在main中被close也会触发这个
			defer close(buf.pending)
			for data := range buf.pending {
				buf.rawSend(data)
			}
			// close connection
			buf.conn.Close()
			fmt.Println("Close connection")
			return
		}
	}
}

// send packet
// !IMPORTANT! once closed, never Send!!!
func (buf *SenderBuffer) Send(data []byte) (err error) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Buffer.Send failed", x)
		}
	}()

	// TODO: 使用Session的Encoder来处理发送的数据(加密等)
	// if buf.sess.Flag&SESS_ENCRYPT != 0 { // if encryption has setup
	// 	buf.sess.Encoder.Codec(data)
	// } else if buf.sess.Flag&SESS_KEYEXCG != 0 { // whether we just exchanged the key
	// 	buf.sess.Flag &= ^SESS_KEYEXCG
	// 	buf.sess.Flag |= SESS_ENCRYPT
	// }

	buf.pending <- data
	return nil
}

// packet online
func (buf *SenderBuffer) rawSend(data []byte) {
	// TODO: 最后又封包一次,加上包头等, 其实可以在上面加密那次封包做了
	// writer := packet.Writer()
	// writer.WriteU16(uint16(len(data)))
	// writer.WriteRawBytes(data)

	// //nr := int16(data[0])<<8 | int16(data[1])
	// //log.Printf("\033[37;44m[ACK] %v\t%v\tSIZE:%v\033[0m\n", nr, proto.RCode[nr], len(data))
	// n, err := buf.conn.Write(writer.Data())
	n, err := buf.conn.Write(data)
	if err != nil {
		fmt.Println("Error send reply, bytes:", n, "reason:", err)
		return
	}
}

// create a new write buffer
func NewSenderBuffer(sess *Session, conn net.Conn, ctrl chan bool) *SenderBuffer {
	max := DEFAULT_OUTQUEUE_SIZE
	buf := SenderBuffer{conn: conn}
	buf.sess = sess
	buf.pending = make(chan []byte, max)
	buf.ctrl = ctrl
	buf.max = max
	return &buf
}
