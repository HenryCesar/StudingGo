package models

import "time"

type Fibonacci struct {
	Input    uint64        `json:"input"`
	Result   uint64        `json:"result"`
	Duration time.Duration `json:"duration"`
}
