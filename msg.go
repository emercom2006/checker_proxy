package main

import (
	"log"
	"strings"
)

func msgtg(msg string) {
	var msgip string
	var urltg string

	urltg = "https://api.telegram.org/bot828966250:AAHW04-O3KVVab_zKdiZEzkKZU3p5uYZmLc/sendMessage?chat_id=132592606&text="
	msgip = urltg + msg

	//Забираем из файла proxytg.txt прокси HTTP протокола для обхода РКН
	mapProxy := ReadFile("proxytg.txt")

	for _, proxy := range mapProxy {
		sliceproxy := strings.Split(proxy, ":")
		ip := sliceproxy[0]
		port := sliceproxy[1]
		user := sliceproxy[2]
		pass := sliceproxy[3]

		proxy = "http://" + user + ":" + pass + "@" + ip + ":" + port

		client := transportProxy(ip, port, user, pass)

		// Шлем запрос по прокси в google
		_, err := client.Get(msgip)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
