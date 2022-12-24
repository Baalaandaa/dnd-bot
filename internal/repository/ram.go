package repository

import (
	"dnd-bot/internal/models"
)

type RAMRepository struct {
	cubeStates    map[int64]models.CubeState
	formulaStates map[int64]models.FormulaState
}

func (R RAMRepository) GetCubeState(user int64) models.CubeState {
	if state, ok := R.cubeStates[user]; ok {
		return state
	}
	R.cubeStates[user] = models.DefaultCubeState
	return models.DefaultCubeState
}

func (R RAMRepository) UpdateCubeState(user int64, state models.CubeState) {
	R.cubeStates[user] = state
}

func (R RAMRepository) GetFormulaState(user int64) models.FormulaState {
	if state, ok := R.formulaStates[user]; ok {
		return state
	}
	R.formulaStates[user] = models.DefaultFormulaState
	return models.DefaultFormulaState
}

func (R RAMRepository) UpdateFormulaState(user int64, state models.FormulaState) {
	R.formulaStates[user] = state
}

func NewRAMRepository() Repository {
	return &RAMRepository{
		cubeStates:    make(map[int64]models.CubeState),
		formulaStates: make(map[int64]models.FormulaState),
	}
}
