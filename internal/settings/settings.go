package settings

import (
	"bytes"
	"os"
	"path/filepath"
	"proxclient/internal/logging"

	"github.com/sirupsen/logrus"
)

const JSON_FILENAME = "prox_config.txt"
const JSON_FILEPATH = "C:\\Windows\\system32"
const PASSPHRASE = "proxccp"

func init() {

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

// Save settings on prox_client.txt file in the system32 directory
func (c *Configs) Save() error {

	config := c.toByte()
	text, _ := c.encrypt(config.Bytes())

	err := os.WriteFile(filepath.Join(JSON_FILEPATH, JSON_FILENAME), text, 0644)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't write to file")
		return err
	}
	return nil
}

// Create a prox_client.txt file in the system32 directory
func (c *Configs) Create() error {
	_, err := os.OpenFile(JSON_FILENAME, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't open file")
		return err
	}

	return nil
}

// Add an client account to settings struct
func (c *Configs) AddAccount(account Account) {
	c.AuthorizedAccounts = append(c.AuthorizedAccounts, account)
}

// Get settings from prox_client.txt file in the system32 directory
func (c *Configs) RecieveStoragedData() error {
	data, err := os.ReadFile(filepath.Join(JSON_FILEPATH, JSON_FILENAME))
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't read file")
		return err
	}

	text, err := c.decrypt(data)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't decrypt file")
		return err
	}
	byteBuffer := bytes.NewBuffer([]byte(text))
	c.toStruct(*byteBuffer)

	return nil
}

// Checks if the client is registered as an authorized
func (cfg *Configs) Contains(account Account) bool {
	for _, c := range cfg.AuthorizedAccounts {
		if c.Name == account.Name {
			return true
		}
	}
	return false
}

// Checks if the is a prox_client.txt file system32 directory
func ValidateConfigs() bool {
	info, err := os.Stat(filepath.Join(JSON_FILEPATH, JSON_FILENAME))
	if os.IsNotExist(err) {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Coundn't validate file status")
		return false
	}
	return !info.IsDir()
}
