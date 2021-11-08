package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/henRok1/telegram-bot/pkg/repository"
)

func(b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error{
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	if err != nil{
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID,
		fmt.Sprintf(b.messages.Start, authLink))
	_, err = b.bot.Send(msg)
	return err
}

func(b *Bot) getAccessToken(chatID int64) (string, error){
	return b.tokenRepository.Get(chatID, repository.AccessTokens)
}


func (b *Bot) generateAuthorizationLink(chadID int64) (string, error){
	redirectURL := b.generateRedicerctURL(chadID)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil{
		return "", err

	}
	if err := b.tokenRepository.Save(chadID, requestToken, repository.RequestTokens); err != nil{
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (b *Bot) generateRedicerctURL(chatID int64) string{
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}