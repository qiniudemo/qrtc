package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"os"
	
	"github.com/gin-gonic/gin"
	"pili-sdk-go.v2/pili"
)

//创建房间
/*
POST /v1/rooms
Host: rtc.qiniuapi.com
Authorization: <QiniuToken>
Content-Type: application/json
{
"owner_id": "<OwnerUserId>",
"room_name": "<RoomName>",
"user_max": "<UserMax>"
}
*/

type Room struct {
	OwnerID  string `json:"owner_id"`
	RoomName string `json:"room_name,omitempty"`
	UserMax  int    `json:"user_max,omitempty"`
}

type createRoomResp struct {
	RoomName string `json:"room_name,omitempty"`
	Error    string `json:"error,omitempty"`
}

func GetCreateRoom(c *gin.Context) {
	if !CheckMac() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "create.tmpl", gin.H{
		"title": "创建房间",
	})
}

func PostCreateRoom(c *gin.Context) {

	maxUser := c.PostForm("user_max")
	var maxUserNum int
	if maxUser != "" {
		num, atoiErr := strconv.Atoi(maxUser)
		if atoiErr != nil {
			c.Error(atoiErr)
		}
		maxUserNum = num
	}

	creator := Room{
		OwnerID:  c.PostForm("owner_id"),
		RoomName: c.PostForm("room_name"),
		UserMax:  maxUserNum,
	}

	reqData, marshErr := json.Marshal(creator)
	if marshErr != nil {
		c.Error(marshErr)

	}

	var createRet createRoomResp

	client := pili.New(Mac, nil)
	callErr := client.CallWithJSON(&createRet, "POST", "http://rtc.qiniuapi.com/v1/rooms", creator)
	if callErr != nil {
		log.Println(callErr)
	}

	if createRet.Error != "" {
		c.String(http.StatusBadRequest, "error: %s", createRet.Error)
	} else {
		DB.Put([]byte(c.PostForm("room_name")), reqData, nil)
		c.Redirect(http.StatusFound, "/")
	}

}

type RoomStatus struct {
	RoomName   string      `json:"room_name,omitempty"`
	OwnerID    string      `json:"owner_id,omitempty"`
	RoomStatus interface{} `json:"room_status,omitempty"`
	UserMax    int         `json:"user_max,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func GetCheckRoom(c *gin.Context) {
	if !CheckMac() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "room.tmpl", gin.H{
		"title": "查看房间",
	})
}

func PostCheckRoom(c *gin.Context) {

	roomName := c.PostForm("room_name")
	reqPath := fmt.Sprintf("http://rtc.qiniuapi.com/v1/rooms/%s", roomName)

	var roomStat RoomStatus
	client := pili.New(Mac, nil)
	reqErr := client.Call(&roomStat, "GET", reqPath)
	if reqErr != nil {
		log.Println("err", reqErr, roomStat.Error)
	}

	RoomStatInJson, umErr := json.Marshal(roomStat)
	if umErr != nil {
		c.Error(umErr)
	}
	c.String(http.StatusOK, string(RoomStatInJson))

}

type deleteResp struct {
	Error string `json:"error"`
}

func GetDeleteRoom(c *gin.Context) {
	if !CheckMac() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "delete.tmpl", gin.H{
		"title": "删除房间",
	})
}

func PostDeleteRoom(c *gin.Context) {
	roomName := c.PostForm("room_name")
	reqPath := fmt.Sprintf("http://rtc.qiniuapi.com/v1/rooms/%s", roomName)

	var delStatus deleteResp

	client := pili.New(Mac, nil)
	reqErr := client.Call(&delStatus, "DELETE", reqPath)
	if reqErr != nil {
		c.Error(reqErr)
	}
	if delStatus.Error != "" {
		c.String(http.StatusBadRequest, "error: %s", delStatus.Error)
	} else {
		DB.Delete([]byte(c.PostForm("room_name")), nil)
		c.Redirect(http.StatusFound, "/delete")
	}
}

func GetToken(c *gin.Context) {
	if !CheckMac() {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	c.HTML(http.StatusOK, "token.tmpl", gin.H{
		"title": "token",
	})
}

func GetAllRooms(c *gin.Context) {

	Rooms := make(map[string]Room)
	iter := DB.NewIterator(nil, nil)
	var tempRoom Room

	for iter.Next() {
		json.Unmarshal(iter.Value(), &tempRoom)
		Rooms[string(iter.Key())] = tempRoom
	}
	RoomsInJson, umErr := json.Marshal(Rooms)
	if umErr != nil {
		c.Error(umErr)
	}
	//c.JSON(http.StatusOK, RoomsInJson)
	c.String(http.StatusOK, string(RoomsInJson))
}

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"title": "auth",
	})
}

func PostLogin(c *gin.Context) {
	ak := c.PostForm("ak")
	sk := c.PostForm("sk")
	auth := Config{
		AK: ak,
		SK: sk,
	}
	Mac.AccessKey = ak
	Mac.SecretKey = []byte(sk)

	confFile, err := os.OpenFile("conf.json",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer confFile.Close()
	data,_ := json.Marshal(auth)
	_,wErr:=confFile.Write(data)
	if wErr!=nil{
		fmt.Println(wErr)
	}
	
	c.Redirect(http.StatusFound, "/")
}

func CheckMac() bool {
	
	if (len(Mac.SecretKey)+len(Mac.AccessKey)) == 0 {
		return false
	}
	return true
}
