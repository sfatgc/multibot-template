package multibot

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

type TgBotBulgakteer struct {
	TgBot
}

func (b *TgBotBulgakteer) Handle(c telebot.Context) error {
	message := c.Text()
	response := fmt.Sprintf("Bulgakteer bot says: %s", message)
	_, err := c.Bot().Send(c.Sender(), response)
	return err
}
