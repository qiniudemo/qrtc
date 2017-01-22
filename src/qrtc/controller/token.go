package controller

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TokenGenerator(c *gin.Context) {

	expireStr := c.PostForm("expire_at")
	expireTime, _ := strconv.Atoi((expireStr))
	expireTime64 := int64(expireTime)
	token, err := CreateToken(c.PostForm("room_name"), c.PostForm("user_id"), c.PostForm("perm"), expireTime64)

	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, "%s", token)
}

func TokenAPI(c *gin.Context) {

	token, err := CreatePermanentToken(c.Param("room_name"), c.Param("user_id"), "user")
	if err != nil {
		c.Error(err)
	}
	c.String(http.StatusOK, "%s", token)

}

type RoomAccess struct {
	RoomName string `json:"room_name"`
	UserID   string `json:"user_id"`
	Perm     string `json:"perm"`
	ExpireAt int64  `json:"expire_at,omitempty"`
}

func CreateToken(roomName, UserID, perm string, expireTime int64) (ret string, err error) {

	rAccess := RoomAccess{
		RoomName: roomName,
		UserID:   UserID,
		Perm:     perm,
		ExpireAt: expireTime,
	}
	RoomAccess, marErr := json.Marshal(rAccess)
	if marErr != nil {
		err = marErr
		return
	}
	encodedRoomAccess := base64.URLEncoding.EncodeToString(RoomAccess)
	h := hmac.New(sha1.New, Mac.SecretKey)
	h.Write([]byte(encodedRoomAccess))

	encodedSign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	ret = fmt.Sprintf("%s:%s:%s", Mac.AccessKey, encodedSign, encodedRoomAccess)

	return
}

func CreatePermanentToken(roomName, UserID, perm string) (ret string, err error) {

	rAccess := RoomAccess{
		RoomName: roomName,
		UserID:   UserID,
		Perm:     perm,
	}
	RoomAccess, marErr := json.Marshal(rAccess)
	if marErr != nil {
		err = marErr
		return
	}
	encodedRoomAccess := base64.URLEncoding.EncodeToString(RoomAccess)
	h := hmac.New(sha1.New, Mac.SecretKey)
	h.Write([]byte(encodedRoomAccess))

	encodedSign := base64.URLEncoding.EncodeToString(h.Sum(nil))
	ret = fmt.Sprintf("%s:%s:%s", Mac.AccessKey, encodedSign, encodedRoomAccess)

	return
}
