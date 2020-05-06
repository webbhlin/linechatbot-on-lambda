package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/webbhlin/awstranslate"
)

var bot *linebot.Client

func getLineEvents(c *gin.Context) ([]*linebot.Event, error) {
	//use another way to get events instead using bot.ParseRequest()

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	request := &struct {
		Events []*linebot.Event `json:"events"`
	}{}
	if err = json.Unmarshal(body, request); err != nil {
		return nil, err
	}
	return request.Events, nil
}

func webhookHandler(c *gin.Context) {
	events, err := getLineEvents(c)
	if err != nil {
		log.Printf("error: %v", err)
	}
	for _, event := range events {
		log.Printf("Event ReplyToken: %v", event.ReplyToken)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "@pahuddev" {
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我猜你你喜歡虎哥, 請看下面 - https://www.youtube.com/watch?v=ziBpmUZc9YQ&t=94s")).Do()
					if err != nil {
						log.Printf("error: %v", err)
					}
				} else if message.Text == "@@" {
					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("AWS chatbot demo, commands: @pahuddev, 2tw, 2en")).Do()
					if err != nil {
						log.Printf("@@error: %v", err)
					}
				} else if strings.Contains(message.Text, "2tw") {
					cmd := strings.Fields(message.Text)
					trText := strings.Join(append(cmd[1:]), " ")

					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(awstranslate.TranslateTo(trText, "en", "zh-tw", "us-east-1"))).Do()
					if err != nil {
						log.Print(err)
					}
				} else if strings.Contains(message.Text, "2en") {
					cmd := strings.Fields(message.Text)
					trText := strings.Join(append(cmd[1:]), " ")

					_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(awstranslate.TranslateTo(trText, "zh-tw", "en", "us-east-1"))).Do()
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/hello", webhookHandler)

	return r
}

func main() {
	//addr := ":" + os.Getenv("PORT")
	channelSecret := os.Getenv("CSECRET")
	botToken := os.Getenv("BTOKEN")
	bot, _ = linebot.New(channelSecret, botToken)

	addr := ":8888" // I don't really use it :P
	log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}
