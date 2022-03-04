package controllers

import (
	"api-rest/models"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	handler   time.Duration
	all       = []*models.Fibonacci{}
	check     *models.Fibonacci
	fibo      *models.Fibonacci
	ok        int
	cont      int
	fibResult uint64
	mutex     sync.RWMutex
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	handler = time.Duration(elapsed.Seconds())
}

func fibonacciCaller(x uint64) {
	defer timeTrack(time.Now())
	fibResult = callFibonacci(x)

	mutex.Lock()
	cont++
	fibo = &models.Fibonacci{
		Input:    x,
		Result:   fibResult,
		Duration: handler,
	}
	all = append(all, fibo)
	mutex.Unlock()
}

func callFibonacci(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return callFibonacci(x-1) + callFibonacci(x-2)
}

func checkFibonacci(input uint64) {
	for _, c := range all {
		if c.Input == input {
			check = c
			ok = 1
			break
		}
	}
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
			"done": false,
		})
	}
	checkFibonacci(input)

	if ok == 1 {
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"done": true,
			"fib": fiber.Map{
				"all": check,
			},
		})
	} else {
		go fibonacciCaller(input)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"done": false,
		})
	}
}
