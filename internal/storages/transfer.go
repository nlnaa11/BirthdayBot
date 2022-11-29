package storages

import (
	"encoding/json"
	"os"

	"github.com/nlnaa11/BirthdayBot/internal/model/users"
	"github.com/pkg/errors"
)

type ChatInfoRecord struct {
	ID       int64  `json:"ID"`
	Forms    string `json:"Link"`
	Database string `json:"Table"`
}

type AddedRecord struct {
	Added []users.User `json:"Added"`
}

type RemovedRecord struct {
	Removed []users.User `json:"Removed"`
}

type BirthdayRecord struct {
	Birthday []users.User `json:"Birthday"`
}

// установить в качестве ключа chatID
func NewTransferFile(path string, record *ChatInfoRecord) error {
	if path == "" {
		return errors.New("invalid path: " + path)
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		return err
	}
	defer f.Close()

	if record == nil {
		return errors.New("invalid record")
	}

	encodeJson := json.NewEncoder(f)
	err = encodeJson.Encode(*record)
	if err != nil {
		return errors.Wrap(err, "adding a chatId to a transfer file")
	}

	_, err = f.WriteString("\n")
	if err != nil {
		return errors.Wrap(err, "adding an empty string")
	}

	return nil
}

func checkChatID(path string, chatID int64) (f *os.File, err error) {
	f, err = os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0555)
	if err != nil {
		return nil, errors.Wrap(err, "opening a transfer file: ")
	}
	//defer f.Close()

	var info ChatInfoRecord
	decodeJson := json.NewDecoder(f)
	err = decodeJson.Decode(&info)
	if err != nil {
		return nil, errors.Wrap(err, "reading a chat info from a file")
	}

	key := info.ID
	if key != chatID {
		return nil, errors.New("this file belongs to another chat")
	}

	return
}

func AddedEntry(path string, chatID int64, record *AddedRecord) error {
	f, err := checkChatID(path, chatID)
	if err != nil {
		return err
	}
	if f == nil {
		return errors.New("a file pointer is lost")
	}
	defer f.Close()

	if record == nil {
		return errors.New("invalid record")
	}

	encodeJson := json.NewEncoder(f)
	err = encodeJson.Encode(*record)
	if err != nil {
		return errors.Wrap(err, "adding a new users: ")
	}

	return nil
}

func RemovedEntry(path string, chatID int64, record *RemovedRecord) error {
	f, err := checkChatID(path, chatID)
	if err != nil {
		return err
	}
	if f == nil {
		return errors.New("a file pointer is lost")
	}
	defer f.Close()

	if record == nil {
		return errors.New("invalid record")
	}

	encodeJson := json.NewEncoder(f)
	err = encodeJson.Encode(*record)
	if err != nil {
		return errors.Wrap(err, "adding an old user")
	}

	return nil
}

func CheckBirthdays(path string, chatID int64) (*BirthdayRecord, error) {
	f, err := checkChatID(path, chatID)
	if err != nil {
		return nil, err
	}
	if f == nil {
		return nil, errors.New("a file pointer is lost")
	}
	defer f.Close()

	var birthdays BirthdayRecord

	decoderJson := json.NewDecoder(f)
	err = decoderJson.Decode(&birthdays)
	if err != nil {
		return nil, errors.Wrap(err, "getting birthdays")
	}

	return &birthdays, nil
}
