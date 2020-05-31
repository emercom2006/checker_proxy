package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func transportProxy(ip, port, user, pass string) *http.Client {

	var proxy = "http://" + user + ":" + pass + "@" + ip + ":" + port

	//создаем url прокси
	var proxyURL, err = url.Parse(proxy)

	if err != nil {
		log.Println(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	//Настройка http Client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 60,
	}
	return client
}
