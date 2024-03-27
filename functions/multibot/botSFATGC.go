package multibot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

type TgBotSFATGC struct {
	TgBot
}

func (b *TgBotSFATGC) Handle(c telebot.Context) error {
	message := c.Text()
	response := fmt.Sprintf("SFATGC bot says: I am %s\n You asked for %s", b.BotName, message)
	_, err := c.Bot().Send(c.Sender(), response)
	return err
}
