package main

import (
	log "github.com/sirupsen/logrus"
	"myapp/bot"
	"myapp/config"
)

//"myapp/bot" //we will create this later
//"myapp/config" //we will create this later

func main() {
	log.Info() // just to get the compiler off my back for now
	err := config.ReadConfig()
	if err != nil {
		log.WithError(err).Error()
		return
	}

	bot.Start()

	// Not gonna lie, I'm just assuming the guide knows
	// 		what it's talking about and that this is
	//		necessary. Will remove if not I guess
	<-make(chan struct{})

	return
}
