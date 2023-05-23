package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:65432")

	if err != nil {
		fmt.Println(err)
	}

	_, err = conn.Write([]byte("1+2,3+4,5+6,7+8,8+9"))

	if err != nil {
		fmt.Println(err)
	}

	

	sumStream := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	n, err := reader.Read(sumStream[:])
	if err != nil {
		fmt.Println(err)
	}
	sumStr := string(sumStream[:n])

	sum, _ := strconv.Atoi(sumStr)
	

	fmt.Println("Total sum:", sum)
}
