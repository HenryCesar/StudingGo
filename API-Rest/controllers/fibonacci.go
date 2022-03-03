package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Fibonacci struct {
	Input    uint64        `json:"input"`
	Result   uint64        `json:"result"`
	Duration time.Duration `json:"duration"`
}

var all = []*Fibonacci{}
var handler time.Duration

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	handler = time.Duration(elapsed.Seconds())
}

func fibonacciCaller(x uint64) uint64 {
	defer timeTrack(time.Now())
	y := callFibonacci(x)
	return y
}

func callFibonacci(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return callFibonacci(x-1) + callFibonacci(x-2)
}

func GetAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"queries": all,
	})
}

func GetNumber(c *fiber.Ctx) error {
	body := c.Params("input")
	input, err := strconv.ParseUint(body, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Input",
		})
	}

	var check *Fibonacci
	var ok int
	for _, c := range all {
		if c.Input == input {
			check = c
			ok = 1
			break
		}
	}

	var fibResult uint64
	var fibo *Fibonacci
	if ok == 1 {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"done": true,
			"fib": fiber.Map{
				"all": check,
			},
		})
	} else {
		fibResult = fibonacciCaller(input)
		fibo = &Fibonacci{
			Input:    input,
			Result:   fibResult,
			Duration: handler,
		}
		all = append(all, fibo)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"done": true,
			"fib": fiber.Map{
				"all": fibo,
			},
		})
	}

}
