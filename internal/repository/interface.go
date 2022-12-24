package repository

import (
	"dnd-bot/internal/models"
)

type Repository interface {
	GetCubeState(user int64) models.CubeState
	UpdateCubeState(user int64, state models.CubeState)

	GetFormulaState(user int64) models.FormulaState
	UpdateFormulaState(user int64, state models.FormulaState)
}
