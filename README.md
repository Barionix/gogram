# Telego 0.1
### A Wrapper to the telegram bot API

[![GoDoc](https://godoc.org/github.com/JuniorMario/gogram?status.svg)](http://godoc.org/github.com/JuniorMario/gogram)
[![Travis](https://travis-ci.org/JuniorMario/gogram.svg?branch=master)](https://travis-ci.org/github/JuniorMario/gogram)

The Telego package makes the API handling easier so you can easily build a telegram bot with a few lines of code.

## Installing
You can install the package by using the `go get` command.
` go get -u github.com/JuniorMario/gogram` 

or you can clone the repositoy and move it to the your GOPATH
``` 
$ git clone https://github.com/JuniorMario/gogram
$ mv telego $GOPATH
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


## Colaborators
| Title | Telegram | Github | Linkedin |
| ------ | ------ | ------ | ------ |
| Developer | [Mario](t.me/Barionixx) | [JuniorMario](https://github.com/JuniorMario/) | [Mário Soares](https://www.linkedin.com/in/mariojrsoares/) |

