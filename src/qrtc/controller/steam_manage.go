package controller

import (
	"net/http"
	
	"pili-sdk-go.v2/pili"
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/json"
)

func GetPublishUrl(c *gin.Context) {
	c.HTML(http.StatusOK, "stream.tmpl", gin.H{
		"title": "stream",
	})
}

func PostPublisUrl(c *gin.Context) {

	domain := c.PostForm("domain")
	hub := c.PostForm("hub")

	stream := StreamManager{
		StreamDomain:domain,
		HubName:hub,
	}
	
	str,err := json.Marshal(stream)
	if err!=nil{
		fmt.Println(err)
	}
	
	StreamDB.Put([]byte("stream"),str,nil)

	c.Redirect(http.StatusFound, "/")

}


func GenPubUrl(c *gin.Context){
	
	stream,_ := StreamDB.Get([]byte("stream"),nil)
	
	var strMan StreamManager
	json.Unmarshal(stream,&strMan)
	
	pubDomain := fmt.Sprintf("pili-publish.%s",strMan.StreamDomain)
	addr := pili.RTMPPublishURL(pubDomain,strMan.HubName,c.Param("room_name"),Mac,43200)
	c.String(http.StatusOK,"%s",addr)
}


func GenPlayUrl(c *gin.Context){
	stream,_ := StreamDB.Get([]byte("stream"),nil)
	
	var strMan StreamManager
	json.Unmarshal(stream,&strMan)
	
	playDomain := fmt.Sprintf("pili-live-rtmp.%s",strMan.StreamDomain)
	addr := pili.RTMPPlayURL(playDomain,strMan.HubName,c.Param("room_name"))
	c.String(http.StatusOK,"%s",addr)
}