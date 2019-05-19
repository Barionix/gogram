package main 

import (
	"fmt"
	"telego"

)

func main() { 
	bot := telego.Conf("821567778:AAGIyRaigJbzdTl9V4Jl8M3fWBCvbiqaTrk")
	a := bot.SetWebhookWithCert("https://webhook.site/29cc0d73-c804-4bf4-bb1e-fb6f83208d45", "cert.pem")
	fmt.Println(a)
	c, d := bot.GetWebhookinfo()
	fmt.Println(c, d)
	b := bot.DeleteWebhook()
	fmt.Println(b)

	
	a = bot.SetWebhook("https://webhook.site/29cc0d73-c804-4bf4-bb1e-fb6f83208d45")
	fmt.Println("res of set", a)
	
	c, d = bot.GetWebhookinfo()
	fmt.Println(c, d)
	b = bot.DeleteWebhook()
	fmt.Println(b)

	
}