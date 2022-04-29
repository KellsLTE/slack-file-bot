package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return os.Getenv(key)
}

func main()  {
	botToken := env("SLACK_BOT_TOKEN")
	channelID := env("SLACK_CHANNEL_ID")
 
	api := slack.New(botToken)
	channelArr := []string{channelID}
	fileArr := []string{"file.txt"}

	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}