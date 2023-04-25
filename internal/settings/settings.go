package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var dir string

const JSON_FILENAME = "prox_config.json"

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

func (c *Configs) Save() {
	config, _ := json.MarshalIndent(c, "", " ")

	err := ioutil.WriteFile(filepath.Join(dir, JSON_FILENAME), config, 0644)
	if err != nil {
		log.Println(err)
	}
}

func (cfg *Configs) AddAccount(account Account) {
	cfg.AuthorizedAccounts = append(cfg.AuthorizedAccounts, account)
}

func (cfg *Configs) RecieveStoragedData() error {
	data, err := ioutil.ReadFile("./prox_config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return err
	}
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
	info, err := os.Stat(filepath.Join(dir, JSON_FILENAME))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
