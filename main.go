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
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gambunbot/code"
	"gambunbot/gacha"
	"gambunbot/osusume"
	"gambunbot/random_pics"
	"gambunbot/utils"

	"github.com/akrylysov/algnhsa"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const (
	INITIAL_VALUE int = 0
	ONE           int = 1
)

func ArrayRand(elements []string) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	randomIndex := INITIAL_VALUE
	if len(elements) > ONE {
		randomIndex = rand.Intn(len(elements))
	}
	return randomIndex
}

func sendReply(bot *linebot.Client, event *linebot.Event, replyMessage []string) {
	index := ArrayRand(replyMessage)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage[index])).Do(); err != nil {
		log.Print(err)
	}
}

func sendReplyAndSticker(bot *linebot.Client, event *linebot.Event, replyMessage []string, pkgSticker, pickSticker string) {
	index := ArrayRand(replyMessage)

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage[index]), linebot.NewStickerMessage(pkgSticker, pickSticker)).Do(); err != nil {
		log.Print(err)
	}
}

func handleCallback(w http.ResponseWriter, req *http.Request, bot *linebot.Client) {
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
				lowerCaseMessage := strings.ToLower(message.Text)
				if strings.Contains(lowerCaseMessage, "#help") || strings.Contains(lowerCaseMessage, "#tasuke") {
					replyMessage := []string{}
					help := utils.Help()
					replyMessage = append(replyMessage, help)
					sendReply(bot, event, replyMessage)
				}

				if strings.Contains(lowerCaseMessage, "#apakah") {
					replyMessage := []string{}
					var pkgSticker, pickSticker string
					if strings.Contains(message.Text, "atau") {
						replyString := []string{}
						if strings.Contains(lowerCaseMessage, "?") {
							splitMessage := regexp.MustCompile(`(\?|\.|!)`).Split(lowerCaseMessage, -1)
							replyString = strings.Split(splitMessage[0], "apakah ")
						} else {
							replyString = strings.Split(lowerCaseMessage, "apakah ")
						}

						replyMessage = strings.Split(replyString[1], " atau ")

					} else if strings.Contains(lowerCaseMessage, "gacha") {
						gachaResult, gachaPercentage := gacha.GachaPercentage()
						replyMessage = append(replyMessage, gachaResult)

						if gachaPercentage >= 75 {
							pkgSticker, pickSticker = gacha.HappyReaction()
						} else if gachaPercentage <= 44 {
							pkgSticker, pickSticker = gacha.SadReaction()
						}
					} else {
						replyMessage = []string{"ya", "tidak", "ya", "tidak"}
					}

					if pkgSticker != "" {
						sendReplyAndSticker(bot, event, replyMessage, pkgSticker, pickSticker)
					}
					sendReply(bot, event, replyMessage)
				} else if strings.Contains(lowerCaseMessage, "#kodenuklir3") || strings.Contains(lowerCaseMessage, "#kodenuklir6") {
					replyMessage := []string{}
					codeResult := code.GetRandomCode(lowerCaseMessage)
					replyMessage = append(replyMessage, codeResult)

					sendReply(bot, event, replyMessage)
				} else if strings.Contains(lowerCaseMessage, "#osusumeanime") || strings.Contains(lowerCaseMessage, "#osusumemanga") || strings.Contains(lowerCaseMessage, "#osusumevn") {
					replyMessage := []string{}
					codeResult := osusume.GetRandomOsusume(lowerCaseMessage)
					replyMessage = append(replyMessage, codeResult)

					sendReply(bot, event, replyMessage)
				}

				if strings.Contains(lowerCaseMessage, "#cats") {
					orgContent, orgPreview, errs := random_pics.GetCats()

					if errs != "" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errs)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(orgContent, orgPreview)).Do(); err != nil {
							log.Print(err)
						}
					}
				}

				if strings.Contains(lowerCaseMessage, "#dogs") {
					orgContent, orgPreview, errs := random_pics.GetDogs()

					if errs != "" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errs)).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(orgContent, orgPreview)).Do(); err != nil {
							log.Print(err)
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
}

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		log.Print(req)
		handleCallback(w, req, bot)
	})

	// lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)

	is_lambda := os.Getenv("IS_LAMBDA")

	if is_lambda == "yes" {
		algnhsa.ListenAndServe(http.DefaultServeMux, nil)
	} else {
		// This is just sample code.
		// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
			log.Fatal(err)
		}
	}
}
