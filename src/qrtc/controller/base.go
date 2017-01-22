package controller

import (
	"encoding/json"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"pili-sdk-go.v2/pili"
)

type Config struct {
	AK string `json:"access_key"`
	SK string `json:"secret_key"`
}

type StreamManager struct {
	StreamDomain string `json:"stream_domain"`
	HubName      string `json:"hub_name"`
}

var (
	account       Config
	Mac           = &pili.MAC{}
	DB            *leveldb.DB
	StreamDB      *leveldb.DB
)

func init() {
	confFile, err := os.Open("conf.json")
	if err != nil {
		log.Println(err.Error())
		return
	}

	decoder := json.NewDecoder(confFile)

	err = decoder.Decode(&account)
	if err != nil {
		log.Println("Failed decode configuration" + err.Error())
		return
	}
	confFile.Close()

	Mac.AccessKey = account.AK
	Mac.SecretKey = []byte(account.SK)

	DB, err = leveldb.OpenFile("model/rooms", nil)
	if err != nil {
		log.Println("Failed decode configuration" + err.Error())
		return
	}
	
	StreamDB,err = leveldb.OpenFile("model/stream",nil)
	if err != nil {
		log.Println("Failed decode configuration" + err.Error())
		return
	}
}
