package messages

import (
	"fmt"

	"github.com/nlnaa11/BirthdayBot/internal/model/users"
	"github.com/pkg/errors"
)

var (
	allGreeting = "Всем привет!\n\nЯ ваш новый друг, " +
		"который не забывает поздравить с днём рождения. Для начала " +
		"заполните [форму](%s).\n(займёт менее 5 минут)"
	oneGreeting = ", добро пожаловать!\n\n" +
		"Не забудь(те) заполнить [форму](%s).\n(займёт менее 5 минут)"
)

type MessageSender interface {
	SendMessage(text string, chatID int64) error
}

type Model struct {
	client MessageSender
	helper *ModelHelper
}

func NewModel(client MessageSender, helper *ModelHelper) *Model {
	return &Model{
		client: client,
		helper: helper,
	}
}

type Message struct {
	Text     string
	NewUsers []users.User
	OldUser  *users.User
}

func (m *Model) IncomingMessage(msg Message, chatName string, chatID int64) error {
	if len(msg.Text) > 0 {
		if msg.Text == "/start" {
			response, err := m.processStart(chatName, chatID)
			if err != nil {
				return errors.Wrap(err, "IncomingMessage: ")
			}
			return m.client.SendMessage(response, chatID)
		}
	}

	if msg.OldUser != nil {
		if err := m.processKicked(msg.OldUser, chatID); err != nil {
			return errors.Wrap(err, "IncomingMessage: ")
		}
	}

	if len(msg.NewUsers) != 0 {
		response, err := m.processAdded(msg.NewUsers, chatID)
		if err != nil {
			return errors.Wrap(err, "IncomingMessage: ")
		}
		err = m.client.SendMessage(response, chatID)
		if err != nil {
			return errors.Wrap(err, "IncomingMessage: ")
		}
	}

	return nil
}

func (m *Model) processStart(chatName string, chatID int64) (string, error) {
	if m.helper.ChatServiced(chatID) {
		return "", errors.New("The chat is already serviced")
	}

	link, err := m.helper.addChat(chatName, chatID)
	if err != nil {
		return "", errors.Wrap(err, "processStart: ")
	}

	return fmt.Sprintf(allGreeting, link), nil
}

func (m *Model) processKicked(user *users.User, chatID int64) error {
	if !m.helper.ChatServiced(chatID) {
		return errors.New("The chat is not serviced")
	}

	return m.helper.removeUser(user, chatID)
}

func (m *Model) processAdded(users []users.User, chatID int64) (string, error) {
	if !m.helper.ChatServiced(chatID) {
		return "", errors.New("The chat is not serviced")
	}

	link, err := m.helper.addUser(users, chatID)
	if err != nil {
		return "", errors.Wrap(err, "processAdded: ")
	}

	names := users[0].Mention()

	if len(users) != 1 {
		for _, user := range users[1:] {
			names += ", " + user.Mention()
		}
	}

	return fmt.Sprintf(names+oneGreeting, link), nil
}
