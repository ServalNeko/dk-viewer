package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	cfg := &tls.Config{}
	conn, err := tls.Dial("tcp", "koukoku.shadan.open.ad.jp:992", cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		i, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		txt := string(buf[:i])
		if utf8.RuneCountInString(txt) > 4 {
			continue
		}
		fmt.Printf("%s", txt)
	}
}
