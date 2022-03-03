package controllers

import (
	"api-rest/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	all     = []*models.Fibonacci{}
	handler time.Duration
)

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

func checkFibonacci(ok int, check *models.Fibonacci, input uint64) (int, *models.Fibonacci) {
	for _, c := range all {
		if c.Input == input {
			check = c
			ok = 1
			break
		}
	}
	if ok == 1 {
		return 1, check
	} else {
		return 0, check
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

	var (
		check     *models.Fibonacci
		fibo      *models.Fibonacci
		ok        int
		fibResult uint64
	)
	res, check := checkFibonacci(ok, check, input)

	if res == 1 {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"done": true,
			"fib": fiber.Map{
				"all": check,
			},
		})
	} else {
		fibResult = fibonacciCaller(input)
		fibo = &models.Fibonacci{
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
