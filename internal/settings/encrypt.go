package settings

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"io"
	"proxclient/internal/logging"

	"github.com/sirupsen/logrus"
)

func createHash() []byte {
	hash := md5.Sum([]byte(PASSPHRASE))
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash[:])
	return dst
}

func setCipher() (cipher.AEAD, error) {
	key := createHash()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

// Encrypt encrypts data using the passphrase.
func (c *Configs) encrypt(data []byte) ([]byte, error) {
	gcm, err := setCipher()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Encrypt Error")
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Encrypt Error")
		return []byte{}, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt decrypts data using the passphrase.
func (c *Configs) decrypt(data []byte) ([]byte, error) {
	gcm, err := setCipher()
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Encrypt Error")
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logging.Log.WithFields(logrus.Fields{
			"err": err,
		}).Error("Encrypt Error")
		return []byte{}, err
	}
	return plaintext, nil
}

func (c *Configs) toByte() bytes.Buffer {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	enc.Encode(c)
	return network
}

func (c *Configs) toStruct(bf bytes.Buffer) {
	dec := gob.NewDecoder(&bf)
	dec.Decode(&c)
}
