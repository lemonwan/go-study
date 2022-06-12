package testing

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	data := []struct{ a, b, c int }{
		{1, 2, 3},
		{4, 5, 9},
		{5, 5, 10},
	}

	for _, v := range data {
		if sum := add(v.a, v.b); sum != v.c {
			t.Errorf("add(%d, %d); got %d; expected %d", v.a, v.b, v.c, sum)
		}
	}
}

func TestReader(t *testing.T) {
	reader := strings.NewReader("ABCDEF")
	newReader := bufio.NewReaderSize(reader, 0)
	// b := make([]byte, 10)
	peek, _ := newReader.Peek(10)
	t.Logf("%d == %q\n", newReader.Buffered(), peek)
	fmt.Println("10")
	fmt.Printf("%#v\n", 10)
	fmt.Printf("%v\n", 10)
	fmt.Printf("%.4g\n", 123.45)
	sprintf := fmt.Sprintf("%.4g\n", 123.45)
	t.Log(sprintf)
	fmt.Printf("%6.2f\n", 123.45)
}

func TestTcpServer(t *testing.T) {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	t.Log("Server ready to read ...")
	for {
		tcp, err := tcpListener.AcceptTCP()
		if err != nil {
			t.Log("accept err：", err)
			continue
		}
		t.Log("A client connected : ", tcp.RemoteAddr().String())
		go tcpPipe(tcp)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ip := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("Disconnected : ", ip)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	i := 0
	for {
		readString, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(readString)
		time.Sleep(time.Second * 3)
		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"
		b := []byte(msg)
		conn.Write(b)
		i++
		if i > 10 {
			break
		}
	}
}

func TestTcpClient(t *testing.T) {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Log("Client connect error ! " + err.Error())
		return
	}
	defer conn.Close()
	t.Log(conn.LocalAddr().String() + " : Client connected!")
	onMessageReceived(conn)
}

func onMessageReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)
		fmt.Println("writing...")
		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
