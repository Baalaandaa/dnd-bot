package models

type CubeState struct {
	CubeCounter int64
}

var DefaultCubeState = CubeState{
	CubeCounter: 1,
}
