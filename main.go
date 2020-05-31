package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	for {

		//Считываем все прокси из файла которые надо проверять
		mapProxy := ReadFile("proxy.txt")

		for _, proxy := range mapProxy {
			sliceproxy := strings.Split(proxy, ":")
			ip := sliceproxy[0]
			port := sliceproxy[1]
			user := sliceproxy[2]
			pass := sliceproxy[3]
			msg := sliceproxy[4]

			client := transportProxy(ip, port, user, pass)

			//создаем GET запрос
			request, err := http.NewRequest("GET", "https://google.ru", nil)
			if err != nil {
				log.Println(err)
			}

			//Получение ответа
			_, err = client.Do(request)
			if err != nil {
				log.Println(err)
				// Шлем сообщение об ошибке в канал tg
				msgtg(msg)
			}
		}
		time.Sleep(300 * time.Second)
	}
	//Далее идет бот телеграм (можно забить команды на удаление и добавление проксей из списка ...)

	/*
		bot, err := tgbotapi.NewBotAPI("84543250:AA7r04-O3VVab_zKdiZEzkKrgfxjytLc")

		if err != nil {
			log.Panic(err)


		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			reply := "Прокси не работает"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)

		}
	*/
}
