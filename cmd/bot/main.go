package main

import (
	"dnd-bot/internal/controllers/cube"
	"dnd-bot/internal/controllers/formula"
	"dnd-bot/internal/repository"
	"fmt"
	"gopkg.in/telebot.v3"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 2 * time.Second},
	}
	fmt.Println(pref.Token)
	fmt.Println("START")
	bot, err := telebot.NewBot(pref)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("??")

	repo := repository.NewRAMRepository()
	cubeC := cube.NewController(repo)
	formulaC := formula.NewController(repo)
	cubeC.Init(bot)
	formulaC.Init(bot)

	bot.Handle("/start", func(ctx telebot.Context) error {
		menu := &telebot.ReplyMarkup{}
		menu.Reply(menu.Row(menu.Text("Cubes"), menu.Text("Formula")))
		ctx.Reply("Hello!", menu)
		return nil
	})
	fmt.Println("START")
	bot.Start()
}
