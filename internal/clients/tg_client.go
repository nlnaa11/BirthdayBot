package clients

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nlnaa11/BirthdayBot/internal/model/messages"
	"github.com/nlnaa11/BirthdayBot/internal/model/users"
	"github.com/pkg/errors"
)

type TokenGetter interface {
	GetToken() string
}

type TgClient struct {
	client *tgbotapi.BotAPI
}

func NewTgClient(tokenGetter TokenGetter) (*TgClient, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.GetToken())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotAPI")
	}

	client.Debug = true

	return &TgClient{
		client: client,
	}, nil
}

func makeMessage(text string, chatID int64) tgbotapi.MessageConfig {
	config := tgbotapi.NewMessage(chatID, text)
	config.ParseMode = tgbotapi.ModeMarkdown

	return config
}

func (tg *TgClient) Congratulate(text string, chatID int64) {
	// Создать сообщение-поздравление
	// Прикрепить кнопку со ссылкой для поздравления (установить видимость)
	// Отправить
}

func (tg *TgClient) SendMessage(text string, chatID int64) error {
	msg := makeMessage(text, chatID)

	if _, err := tg.client.Send(msg); err != nil {
		return errors.Wrap(err, "client sends a message: ")
	}

	return nil
}

func (tg *TgClient) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0) // getting UpdateCondig: offset 0
	u.Timeout = 60

	updates := tg.client.GetUpdatesChan(u) // getting of chan of Update-s

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil { // if we got a message
			log.Printf("[%s: %d] %s", update.Message.From.UserName,
				update.Message.From.ID, update.Message.Text)

			var newUsers []users.User
			var oldUser *users.User

			text := update.Message.Text

			count := len(update.Message.NewChatMembers)
			if count != 0 {
				for _, user := range update.Message.NewChatMembers {
					if user.IsBot {
						continue
					}
					name := user.FirstName
					if user.LastName != "" {
						name += " " + user.LastName
					}
					newUsers = append(newUsers, users.User{
						UserID:   user.ID,
						UserName: user.UserName,
						FullName: name,
					})
				}
			}

			user := update.Message.LeftChatMember
			if user != nil && !user.IsBot {
				name := user.FirstName
				if user.LastName != "" {
					name += " " + user.LastName
				}

				oldUser = &users.User{
					UserID:   user.ID,
					UserName: user.UserName,
					FullName: name,
				}
			}

			if err := msgModel.IncomingMessage(messages.Message{
				Text:     text,
				NewUsers: newUsers,
				OldUser:  oldUser,
			}, update.Message.Chat.Title, update.Message.Chat.ID); err != nil {
				log.Println("error processing meessage: ", err)
			}
		}
	}
}
