package cube

import (
	"dnd-bot/internal/models"
	"dnd-bot/internal/repository"
	"fmt"
	"gopkg.in/telebot.v3"
	"math/rand"
	"strconv"
)

var (
	menu   = &telebot.ReplyMarkup{}
	btnInc = telebot.Btn{
		Unique: "cubeBtnInc",
		Text:   "+",
	}
	btnDec = telebot.Btn{
		Unique: "cubeBtnDec",
		Text:   "-",
	}

	cubes = menu.Row(
		menu.Data("D2", "cubeRoll", "2"),
		menu.Data("D4", "cubeRoll", "4"),
		menu.Data("D6", "cubeRoll", "6"),
		menu.Data("D8", "cubeRoll", "8"),
		menu.Data("D10", "cubeRoll", "10"),
		menu.Data("D12", "cubeRoll", "12"),
		menu.Data("D20", "cubeRoll", "20"),
		menu.Data("D100", "cubeRoll", "100"),
	)
)

type Controller struct {
	repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {

	return &Controller{
		repo: repo,
	}
}

func (c Controller) Send(ctx telebot.Context) {
	user := ctx.Message().Sender.ID
	menu.Inline(menu.Row(btnDec, btnInc), cubes)
	err := ctx.Reply(fmt.Sprint(c.repo.GetCubeState(user).CubeCounter), menu)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}

func (c Controller) Inc(ctx telebot.Context) {
	user := ctx.Message().Sender.ID
	c.repo.UpdateCubeState(user, models.CubeState{
		CubeCounter: c.repo.GetCubeState(user).CubeCounter + 1,
	})
	menu.Inline(menu.Row(btnDec, btnInc), cubes)
	err := ctx.Edit(fmt.Sprint(c.repo.GetCubeState(user).CubeCounter), menu)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}

func (c Controller) Dec(ctx telebot.Context) {
	user := ctx.Message().Sender.ID
	newVal := c.repo.GetCubeState(user).CubeCounter - 1
	if newVal < 1 {
		newVal = 1
	}
	c.repo.UpdateCubeState(user, models.CubeState{
		CubeCounter: newVal,
	})
	menu.Inline(menu.Row(btnDec, btnInc), cubes)
	err := ctx.Edit(fmt.Sprint(c.repo.GetCubeState(user).CubeCounter), menu)
	if err != nil {
		fmt.Println("ERROR", err)
	}
}

func (c Controller) Roll(ctx telebot.Context) {
	user := ctx.Message().Sender.ID
	size, _ := strconv.ParseInt(ctx.Callback().Data, 10, 64)
	menu.Inline()
	cnt := c.repo.GetCubeState(user).CubeCounter
	score := int64(0)
	solution := "("
	for i := int64(0); i < cnt; i++ {
		rng := rand.Int63n(size) + 1
		score += rng
		if size == 20 && rng == 20 {
			solution += "!20!"
		} else {
			solution += fmt.Sprint(rng)
		}
		if i+1 != cnt {
			solution += "+"
		}
	}
	solution += ")"
	err := ctx.Edit(fmt.Sprintf("%dd%d: %s=%d", cnt, size, solution, score), menu)
	c.repo.UpdateCubeState(user, models.CubeState{
		CubeCounter: 1,
	})
	if err != nil {
		fmt.Println("ERROR", err)
	}
}

func (c Controller) Init(bot *telebot.Bot) {
	bot.Handle("Cubes", func(ctx telebot.Context) error {
		c.Send(ctx)
		return nil
	})
	bot.Handle("\fcubeBtnInc", func(ctx telebot.Context) error {
		c.Inc(ctx)
		return nil
	})
	bot.Handle("\fcubeBtnDec", func(ctx telebot.Context) error {
		c.Dec(ctx)
		return nil
	})
	bot.Handle("\fcubeRoll", func(ctx telebot.Context) error {
		c.Roll(ctx)
		return nil
	})
}
