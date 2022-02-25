package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

const address = ""

var ignoreWord = "ยังไม่ใช่คำตอบที่ถูก"
var fruits = []string{"chayote fruit", "kundong", "dekopon", "rose apple", "mamey sapote", "ackee", "agave plant", "bilimbi", "dead man's fingers", "korlan", "charichuelo fruit", "kahikatea", "babaco", "bilimbi", "calamansi", "clementines", "nere", "loquat", "fibrous satinash", "batuan fruit"}

func main() {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 10224)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))

	c := make(map[string]bool)

	// not optimal solution for brute force fruits
	rand.Seed(time.Now().UnixNano())
	for {
		fruit1 := fruits[rand.Intn(len(fruits))]
		fruit2 := fruits[rand.Intn(len(fruits))]
		fruit3 := fruits[rand.Intn(len(fruits))]

		f := fmt.Sprintf("%s-%s-%s", fruit1, fruit2, fruit3)
		if _, ok := c[f]; !ok {
			_, err := conn.Write([]byte(f))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(50 * time.Millisecond)
		}

		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[:n])
		if !strings.Contains(s, ignoreWord) {
			fmt.Println(s)
		}
	}
}
