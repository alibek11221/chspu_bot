package main

import (
	"CHGPU_T_BOT/lib"
	"CHGPU_T_BOT/lib/mainmenu"
	"CHGPU_T_BOT/lib/state"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal(".env BOT_TOKEN not found")
	}

	webHookURL, ok := os.LookupEnv("WEBHOOK_URL")

	if !ok {
		log.Fatal(".env BOT_TOKEN not found")
	}

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sigs := make(chan os.Signal)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan struct{}, 1)

	updates, _ := bot.UpdatesViaWebhook("/bot"+bot.Token(), telego.WithWebhookSet(&telego.SetWebhookParams{
		URL: webHookURL + "/bot" + bot.Token(),
	}))

	bh, _ := th.NewBotHandler(bot, updates, th.WithStopTimeout(time.Second*5))

	state.PopulateFromCache()

	bh.Handle(handleMessage, th.AnyMessage())

	bh.HandleCallbackQuery(handleBackCommand, th.CallbackDataEqual("back"))

	bh.HandleCallbackQuery(handleInnit, th.CallbackDataEqual("start"))

	go func() {
		<-sigs
		fmt.Println("Stopping")
		_ = bot.StopWebhook()
		state.UpdateCache()
		bh.Stop()
		done <- struct{}{}
	}()
	go bh.Start()
	fmt.Println("Start handling")
	go func() {
		_ = bot.StartWebhook("127.0.0.1:8080")
	}()
	<-done
	fmt.Println("Done")
}
func handleMessage(bot *telego.Bot, update telego.Update) {
	if update.Message.Text == "/start" {
		_, _ = mainmenu.SendInitMessage(bot, update)
		return
	}
	_, err := messageHandler(bot, update)
	if err != nil {
		log.Print("Error", err)
	}
}
func handleBackCommand(bot *telego.Bot, query telego.CallbackQuery) {
	update := telego.Update{Message: query.Message}
	_ = bot.DeleteMessage(tu.Delete(tu.ID(update.Message.Chat.ID), update.Message.MessageID))
	_, err := mainmenu.SendWelcomeBackMessage(bot, update)
	if err != nil {
		os.Exit(1)
	}
}
func handleInnit(bot *telego.Bot, query telego.CallbackQuery) {
	update := telego.Update{Message: query.Message}
	_, err := mainmenu.SendInitMessage(bot, update)
	if err != nil {
		os.Exit(1)
	}
}
func messageHandler(bot *telego.Bot, update telego.Update) (*telego.Message, error) {
	sender, ok := lib.MessageMap[update.Message.Text]
	if ok {
		return sender(bot, update)
	}
	return mainmenu.SendErrorMsg(bot, update)
}
