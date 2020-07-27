# Telego 0.1
### A Wrapper to the telegram bot API

[![Travis](https://travis-ci.org/JuniorMario/telego.svg?branch=master)](https://travis-ci.org/github/JuniorMario/telego)

The Telego package makes the API handling easier so you can easily build a telegram bot with a few lines of code.

### Simple bot example

``` 
package main

import (
	"telego"
)

func main() {
	bot := telego.SetNewBot("<TOKEN>")
	for {
		for _, msg := range bot.GetMsgUpdates().Result {
			fmt.Println(msg)
		}
	}
}
```

The code above start the bot, listen to incomin messages and returns a JSON containing the messages info.
