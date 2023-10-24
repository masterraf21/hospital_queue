package main

import "fmt"

var (
	ErrEmptyQueue         = fmt.Errorf("Empty Queue")
	ErrRoundRobinFailed   = fmt.Errorf("Round Robin failed getting alternate gender")
	ErrModeNotImplemented = fmt.Errorf("Mode not implemented")
	ErrItemInQueue        = fmt.Errorf("Item already in queue")
)
