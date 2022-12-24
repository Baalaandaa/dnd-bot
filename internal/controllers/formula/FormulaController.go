package formula

import (
	"dnd-bot/internal/models"
	"dnd-bot/internal/repository"
	"dnd-bot/pkg/solver"
	"fmt"
	"gopkg.in/telebot.v3"
)

type Controller struct {
	repo repository.Repository
}

func NewController(repo repository.Repository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c Controller) Enter(ctx telebot.Context) {
	ctx.Reply("Now enter formula like `3d8+5`:")
	c.repo.UpdateFormulaState(ctx.Message().Sender.ID, models.FormulaState{Entered: true})
}

func (c Controller) OnText(ctx telebot.Context) {
	if c.repo.GetFormulaState(ctx.Message().Sender.ID).Entered {
		result, err := solver.Calc(ctx.Text())
		if err == nil {
			c.repo.UpdateFormulaState(ctx.Message().Sender.ID, models.FormulaState{Entered: false})
			ctx.Reply(fmt.Sprintf("Avg: %d. Roll: %d. Max(cubes): %d", result.ExpectedScore, result.Score, result.MaxScore))
		} else {
			ctx.Reply(fmt.Sprintf("Send valid formula like `3d8+5`(NO SPACES)"))
		}
	}
}

func (c Controller) Init(bot *telebot.Bot) {
	bot.Handle("Formula", func(ctx telebot.Context) error {
		c.Enter(ctx)
		return nil
	})

	bot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		c.OnText(ctx)
		return nil
	})
}
