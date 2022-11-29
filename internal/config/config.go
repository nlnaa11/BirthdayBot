package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	dir          = "data/"
	tokenFile    = dir + "config.yaml"
	servicedFile = dir + "config.json"
	pathToDb     = dir + "congratulation.db"
)

/**
 * path to db: 				data/congratulation.db,
 * table in db: 			uniqueName
 * path to a transfer file: data/uniqueName.json
 * forms link:				Chat::FormsLink
 */

type Chat struct {
	ChatID     int64  `json:"ID"`
	UniqueName string `json:"uniqueName"`
	FormsLink  string `json:"formsLink,omitempty"`
}

type Config struct {
	Token         string         `yaml:"token"`
	ServicedChats map[int64]Chat `json:"chats"`
}

type Service struct {
	config Config
}

func InitConfig() (*Service, error) {
	s := &Service{
		config: Config{
			ServicedChats: make(map[int64]Chat),
		}}

	f, err := os.OpenFile(tokenFile, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.Wrap(err, "opening a config file: "+tokenFile)
	}
	defer f.Close()

	decoderYaml := yaml.NewDecoder(f)
	err = decoderYaml.Decode(&s.config)
	if err != nil {
		return nil, errors.Wrap(err, "reading a config file: "+tokenFile)
	}

	return s, nil
}

func RecoverConfig() (*Service, error) {
	s, err := InitConfig()
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(servicedFile, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.Wrap(err, "reading a config file: "+servicedFile)
	}
	defer f.Close()

	chats := []Chat{}

	decoderJson := json.NewDecoder(f)
	err = decoderJson.Decode(&chats)
	if err == io.EOF {
		return s, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "reading a config file: "+servicedFile)
	}

	for _, data := range chats {
		s.config.ServicedChats[data.ChatID] = data
	}

	return s, nil
}

func (s *Service) GetToken() string {
	return s.config.Token
}

func (s *Service) GetDatabaseFullPath(chatID int64) (string, string, error) {
	if s.config.ServicedChats == nil {
		return "", "", errors.New("There are not serviced chats")
	}
	if _, exist := s.config.ServicedChats[chatID]; !exist {
		return "", "", errors.New("There is no chat with such an ID")
	}

	return pathToDb, s.config.ServicedChats[chatID].UniqueName, nil
}

func (s *Service) GetDatabasePath() string {
	return pathToDb
}

func (s *Service) GetTransferFilePath(chatID int64) (string, error) {
	if s.config.ServicedChats == nil {
		return "", errors.New("There are not serviced chats")
	}
	if _, exist := s.config.ServicedChats[chatID]; !exist {
		return "", errors.New("There is no chat with such an ID")
	}

	return dir + s.config.ServicedChats[chatID].UniqueName + ".json", nil
}

func (s *Service) GetFormsLink(chatID int64) (string, error) {
	if s.config.ServicedChats == nil {
		return "", errors.New("There are not serviced chats")
	}
	if _, exist := s.config.ServicedChats[chatID]; !exist {
		return "", errors.New("There is no chat with such an ID")
	}

	return s.config.ServicedChats[chatID].FormsLink, nil
}

func (s *Service) GetUniqueName(name string) (string, error) {
	name = strings.ToLower(name)

	files, err := os.ReadDir(dir)
	if err != nil {
		return name, err
	}

	uniqId := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), name) {
			uniqId++
		}
	}

	if uniqId == 0 {
		return name, nil
	}

	return fmt.Sprintf(name+"%d", uniqId), nil
}

func (s *Service) GetStartDir() string {
	return dir
}

func (s *Service) IsServiced(chatID int64) bool {
	_, exist := s.config.ServicedChats[chatID]
	return exist
}

func (s *Service) AddChat(name, link string, chatID int64) error {
	if s.config.ServicedChats == nil {
		return errors.New("Failed to save a new chat info")
	}

	chat := Chat{
		ChatID:     chatID,
		UniqueName: name,
		FormsLink:  link,
	}

	s.config.ServicedChats[chatID] = chat

	return s.addChatToConfig()
}

func (s *Service) addChatToConfig() error {
	f, err := os.OpenFile(servicedFile, os.O_CREATE|os.O_WRONLY, 0655)
	if err != nil {
		return errors.Wrap(err, "opening a config file: "+servicedFile)
	}
	defer f.Close()

	tmp := []Chat{}

	for _, chat := range s.config.ServicedChats {
		tmp = append(tmp, chat)
	}

	encoderJson := json.NewEncoder(f)
	encoderJson.SetIndent("", "    ")
	err = encoderJson.Encode(&tmp)
	if err != nil {
		return errors.Wrap(err, "adding a chat to the config file")
	}

	return nil
}
