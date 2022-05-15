// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gambunbot/gacha"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	// bot, err := linebot.New(
	// 	os.Getenv("CHANNEL_SECRET"),
	// 	os.Getenv("CHANNEL_TOKEN"),
	// )

	bot, err := linebot.New(
		"2656fe09b0298ca731ef9d2dee59e954",
		"Rx0nXjgzZ285hrGebaWjASG/I3UK/kougus6c37nel6iUKAGAUjD4mcoVqXpG9zjmYGJzsAzYMfjArS5N39Z/OxY66eQqPcR+CqXUZRcZZO0Uu+/EQp+UU4fWm+KmS7ql4TA6OrjZtyeKyMqFaAe+gdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if strings.Contains(message.Text, "$apakah gacha") {
						if strings.Contains(message.Text, "$draw") && strings.Contains(message.Text, "$rate") {
							splitter := strings.Split(message.Text, "$")
							draw, _ := strconv.Atoi(strings.Split(splitter[2], " ")[1])
							rate, _ := strconv.Atoi(strings.Split(splitter[3], " ")[1])

							luckMessage, luck := gacha.GachaPercentage()
							simMessage := gacha.GachaSim(draw, rate, 1, luck)
							replyMessage := luckMessage + "\n" + simMessage

							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
								log.Print(err)
							}
						} else {
							replyMessage, percentage := gacha.GachaPercentage()
							var pkgSticker, pickSticker string

							if percentage >= 75 {
								pkgSticker, pickSticker = gacha.HappyReaction()
								if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage), linebot.NewStickerMessage(pkgSticker, pickSticker)).Do(); err != nil {
									log.Print(err)
								}
							} else if percentage <= 44 {
								pkgSticker, pickSticker = gacha.SadReaction()
								if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage), linebot.NewStickerMessage(pkgSticker, pickSticker)).Do(); err != nil {
									log.Print(err)
								}
							} else {
								if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
									log.Print(err)
								}
							}
						}
					}

					if strings.Contains(message.Text, "$gacha sim") {
						splitter := strings.Split(message.Text, "$")
						draw, _ := strconv.Atoi(strings.Split(splitter[2], " ")[1])
						rate, _ := strconv.Atoi(strings.Split(splitter[3], " ")[1])

						replyMessage := gacha.GachaSim(draw, rate, 1, 0)

						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
						}
					}

					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					// 	log.Print(err)
					// }

					// case *linebot.StickerMessage:
					// 	replyMessage := fmt.Sprintf(
					// 		"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					// 	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					// 		log.Print(err)
					// 	}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
