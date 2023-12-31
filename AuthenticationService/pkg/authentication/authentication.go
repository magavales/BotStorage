package authentication

import (
	"AuthenticationService/pkg/database"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strings"
)

var salt = "VX9XQ!psm7Qe#PbSC1#Cvup_4CI/TW7d"

type Authentication struct {
}

func (auth *Authentication) VerificationPassword(data string) (string, bool) {
	var (
		db       database.PostgresDB
		chatID   string
		password string
	)
	data = strings.Split(data, " ")[1]
	chatID, password = parseData(data)
	db.Connect()
	tableData, err := db.Access.GetPwd(db.Pool, chatID)
	if err != nil {
		log.Printf("Don't found password for %s", chatID)
		return "", false
	}

	pwd := hex.EncodeToString(pbkdf2.Key([]byte(password), []byte(tableData.Salt), 10000, 32, sha1.New))
	if tableData.Pwd == pwd {
		log.Println("Password is corrected!")
		return chatID, true
	} else {
		log.Println("Password is uncorrected!")
		return "", false
	}
}

func (auth *Authentication) SetPassword(data string) bool {
	var (
		db       database.PostgresDB
		chatID   string
		password string
	)
	data = strings.Split(data, " ")[1]
	chatID, password = parseData(data)
	pwd := hex.EncodeToString(pbkdf2.Key([]byte(password), []byte(salt), 10000, 32, sha1.New))
	db.Connect()
	err := db.Access.SetPwd(db.Pool, chatID, pwd, salt)
	if err != nil {
		log.Printf("Don't found password for %s", chatID)
		return false
	} else {
		return true
	}
}

func parseData(data string) (string, string) {
	var (
		chatID string
		pwd    string
	)
	decodeData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal("error:", err)
	}
	condition := false
	for _, v := range decodeData {
		if string(v) == ":" {
			condition = true
		} else {
			if condition {
				pwd = pwd + string(v)
			} else {
				chatID = chatID + string(v)
			}
		}
	}
	return chatID, pwd
}
