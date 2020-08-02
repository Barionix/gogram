# Gogram - 0.1
### A Wrapper to the telegram bot API

[![GoDoc](https://godoc.org/github.com/JuniorMario/gogram?status.svg)](http://godoc.org/github.com/JuniorMario/gogram)
[![Travis](https://travis-ci.org/JuniorMario/gogram.svg?branch=master)](https://travis-ci.org/github/JuniorMario/gogram)

The Gogram package makes the API handling easier so you can easily build a telegram bot with a few lines of code.

## Installing
You can install the package by using the `go get` command:
` go get -u github.com/JuniorMario/gogram` 

or you can clone the repositoy and move it to  your GOPATH
``` 
$ git clone https://github.com/JuniorMario/gogram
$ mv gogram $GOPATH
``` 

### Simple bot example

``` 
package main

import (
	"github.com/JuniorMario/gogram"
)

func main() {
	bot := gogram.SetNewBot("<TOKEN>")
	for {
		for _, msg := range bot.GetMsgUpdates().Result {
			fmt.Println(msg)
		}
	}
}
```

The code above start the bot, listen to incomin messages and returns a JSON containing the messages info.

## Available Methdos
| Method | Status | Tested |
| ------ | ------ | ------ |
| Send_Message | Done | Yes |
| SendVideoNote | Done | Yes |
| Reply_To | Done | Yes |
| GetMe | Done | Yes |
| FowardMessage | Done | Yes |
| SendPhoto | Done | Yes |
| SendAudio | Done | Yes |
| SendDocument | Done | Yes |
| SendVideo | Done | Yes |
| SendAnimation | Done | Yes |
| SendVoice | Done | Yes |
| SendLocation | Done | Yes |
| GetAllUpdates | Done | Yes |
| GetMsgUpdates | Done | Yes |
| SetChatTitle | Done | Yes |
| SetChatPermissions | Done | Yes |
| SetChatDescription | Done | Yes |
| setChatAdministratorCustomTitle | Done | Yes |
| kickChatMember | Done | Yes |
| restrictChatMember | Done | Yes |




## Colaborators
| Title | Telegram | Github | Linkedin |
| ------ | ------ | ------ | ------ |
| Developer | [Mario](t.me/Barionixx) | [JuniorMario](https://github.com/JuniorMario/) | [Mário Soares](https://www.linkedin.com/in/mariojrsoares/) |



