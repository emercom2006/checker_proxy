# checker_proxy
 Проверка списка прокси из файла на работоспособность и отсылка уведомлений в телеграм о неработоспособных, путем обхода РКН через прокси.

1. В файле msg.go 12 строка
   urltg = "https://api.telegram.org/bot"ключ tg API"/sendMessage?chat_id="id канала"&text="
   Вставить свои значения, пример: urltg = "https://api.telegram.org/bot843556250:AGHF04-OFDFVVab_zKdfethuyfZU3p5uYZmLc/sendMessage?chat_id=122323606&text="

2. В файл proxy.txt вставляем список прокси для проверки в формате:
 
 IP:Port:Login:Pass:Сообщение, которое надо отослать при неработающем прокси

3. В файл proxytg.txt вставляем прокси не РФ, в случае, если не идет отсылка с родного IP адреса в телеграм.
