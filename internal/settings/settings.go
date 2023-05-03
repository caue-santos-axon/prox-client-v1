package settings

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var dir string

const JSON_FILENAME = "prox_config.txt"
const JSON_FILEPATH = "C:\\Windows\\system32"
const PASSPHRASE = "proxccp"

func init() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = workingDir
}

type Configs struct {
	InboundPath        string `json:"inboundPath" `
	BackupPath         string `json:"backupPath" `
	CreatedOn          string `json:"createdOn" `
	UpdatedOn          string `json:"updatedOn" `
	ReceiveReport      bool   `json:"receiveReport"`
	Key                string
	AuthorizedAccounts []Account
}

type Account struct {
	Name       string `json:"name" `
	PortaLogin string `json:"PortalLogin" `
	CreatedOn  string `json:"createdOn" `
	UpdatedOn  string `json:"updatedOn" `
}

func (c *Configs) Save() error {

	config := c.toByte()
	text, _ := c.encrypt(config.Bytes())

	err := ioutil.WriteFile(filepath.Join(JSON_FILEPATH, JSON_FILENAME), text, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Configs) Create() error {
	_, err := os.OpenFile(JSON_FILENAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//log
		return err
	}

	return nil
}

func (c *Configs) AddAccount(account Account) {
	c.AuthorizedAccounts = append(c.AuthorizedAccounts, account)
}

func (c *Configs) RecieveStoragedData() error {
	data, err := ioutil.ReadFile(JSON_FILENAME)
	if err != nil {
		return err
	}

	text, err := c.decrypt(data)
	if err != nil {
		//log
		return err
	}
	byteBuffer := bytes.NewBuffer([]byte(text))
	c.toStruct(*byteBuffer)

	return nil
}

func (cfg *Configs) Contains(account Account) bool {
	for _, c := range cfg.AuthorizedAccounts {
		if c.Name == account.Name {
			return true
		}
	}
	return false
}

func ValidateConfigs() bool {
	info, err := os.Stat(filepath.Join(JSON_FILEPATH, JSON_FILENAME))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
