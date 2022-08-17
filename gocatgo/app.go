package gocatgo

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/vaaleyard/gocatgo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB           *gorm.DB
	Host         string
	Alphabet     string
	AESCipherkey []byte
}

func (app *App) initializeDB() error {
	dsn := "host=host user=postgres password=pwd dbname=dbname port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Run() error {
	err := app.initializeDB()
	if err != nil {
		return err
	}
	log.Println("Connected to the database")

	err = app.DB.AutoMigrate(&models.Pastebin{})
	if err != nil {
		return err
	}
	log.Println("Database migration completed")

	app.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return nil
}

func AESEncrypt(key []byte, message string) (encoded string, err error) {
	text := []byte(message)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, text, nil)), nil
}

func AESDecrypt(key []byte, encryptedMessage string) (decoded string, err error) {
	messageHexDecoded, _ := hex.DecodeString(encryptedMessage)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return
	}

	gcmNonceSize := gcm.NonceSize()
	if len(messageHexDecoded) < gcmNonceSize {
		return
	}

	nonce, messageHexDecoded := messageHexDecoded[:gcmNonceSize], messageHexDecoded[gcmNonceSize:]
	plainText, err := gcm.Open(nil, nonce, messageHexDecoded, nil)
	if err != nil {
		return
	}

	return fmt.Sprintf("%s", plainText), nil
}
