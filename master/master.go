package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:65432")

	if err != nil {
		fmt.Println(err)
	}

	conn, err := listener.Accept()

	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1024)

	n, err := conn.Read(data)
	data = data[:n]

	if err != nil {
		fmt.Println(err)
	}

	equations := strings.Split(string(data), ",")

	cunkSize := len(equations) / 4

	var chunks [][]string
	for i := 0; i < len(equations); i += cunkSize {
		chunks = append(chunks, equations[i:i+cunkSize])
	}
	fmt.Println(chunks)

	sums := make([]int, 0, 4)

	for _, chunk := range chunks {
		addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:65433")
		if err != nil {
			fmt.Println(err)
		}
		slaveConn, err := net.DialTCP("tcp", nil, addr)
		if err != nil {
			fmt.Println(err)
		}
		_, err = slaveConn.Write([]byte(strings.Join(chunk, ",")))
		if err != nil {
			fmt.Println(err)
		}
		sumStream := make([]byte,1024)
		reader := bufio.NewReader(slaveConn)
		n, err = reader.Read(sumStream[:])
		if err != nil {
			fmt.Println(err)
		}
		sumStr := string(sumStream[:n])
	
		sum, _ := strconv.Atoi(sumStr)
		sums = append(sums, sum)
	}

	totalSum := 0

	for _, sum := range sums {
		totalSum += sum
	}

	_, err = conn.Write([]byte(strconv.Itoa(totalSum)))

	if err != nil {
		fmt.Println(err)
	}
}
