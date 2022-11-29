package messages

import (
	"fmt"

	"github.com/nlnaa11/BirthdayBot/internal/model/users"
	"github.com/nlnaa11/BirthdayBot/internal/storages"
	"github.com/pkg/errors"
)

type ChatsInfo interface {
	GetFormsLink(chatID int64) (string, error)
	GetTransferFilePath(chatID int64) (string, error)
	GetDatabaseFullPath(chatID int64) (string, string, error)
	GetDatabasePath() string
	GetUniqueName(name string) (string, error)
	GetStartDir() string

	IsServiced(chatID int64) bool
	AddChat(uniqueName, formsLink string, chatID int64) error
}

type ModelHelper struct {
	info ChatsInfo
}

func NewModelHelper(info ChatsInfo) *ModelHelper {
	return &ModelHelper{
		info: info,
	}
}

func (h *ModelHelper) ChatServiced(chatID int64) bool {
	return h.info.IsServiced(chatID)
}

func (h *ModelHelper) addChat(chatName string, chatID int64) (string, error) {
	name, err := h.info.GetUniqueName(chatName)
	if err != nil {
		return "", err
	}

	link, err := storages.NewForms(name)
	if err != nil {
		return "", err
	}

	dbPath := h.info.GetDatabasePath()
	err = storages.NewTable(dbPath, name)
	if err != nil {
		return "", err
	}

	dir := h.info.GetStartDir()
	transferPath := fmt.Sprintf(dir + name + ".json")
	data := storages.ChatInfoRecord{
		ID:       chatID,
		Forms:    link,
		Database: name,
	}
	err = storages.NewTransferFile(transferPath, &data)
	if err != nil {
		return "", err
	}

	err = h.info.AddChat(name, link, chatID)
	if err != nil {
		return "", err
	}

	link = "https://pkg.go.dev/time#pkg-types"

	return link, nil
}

func (h *ModelHelper) removeUser(user *users.User, chatID int64) error {
	if user == nil {
		return errors.New("invalid an user info")
	}

	path, err := h.info.GetTransferFilePath(chatID)
	if err != nil {
		return err
	}

	data := storages.RemovedRecord{
		Removed: []users.User{*user},
	}
	return storages.RemovedEntry(path, chatID, &data)
}

func (h *ModelHelper) addUser(users []users.User, chatID int64) (string, error) {
	if len(users) == 0 {
		return "", errors.New("Thare are not new users")
	}

	path, err := h.info.GetTransferFilePath(chatID)
	if err != nil {
		return "", err
	}

	data := storages.AddedRecord{
		Added: users,
	}
	err = storages.AddedEntry(path, chatID, &data)
	if err != nil {
		return "", err
	}

	link, err := h.info.GetFormsLink(chatID)
	if err != nil {
		return "", err
	}

	return link, err
}
