package models

type State int

const (
	ResponseState State = iota
	DoneState
	ErrState
)
