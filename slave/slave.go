package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:65433")

	if err != nil {
		fmt.Println(err)
	}
	for {
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
		fmt.Println(equations)
		sum := 0.0

		for _, eq := range equations {
			expression, _ := govaluate.NewEvaluableExpression(eq)
			resInterface, _ := expression.Evaluate(nil)
			res, ok := resInterface.(float64)
			if !ok {
				fmt.Println("err in evaluate the expression")
			}
			sum += res
		}
		fmt.Println(sum)
		_, _ = conn.Write([]byte(strconv.Itoa(int(sum))))

		conn.Close()
	}
}
