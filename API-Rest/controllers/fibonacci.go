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
	fibResult uint64
	count     int
	completed int
	mutex     sync.RWMutex
)

// Função que recebe o tempo inicial e insere o tempo final na variável para depois converter em segundos.
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	handler = time.Duration(elapsed.Seconds())
}

// Função que chama o Fibonacci
func fibonacciCaller(x uint64) {
	defer timeTrack(time.Now())
	fibResult = callFibonacci(x)

	mutex.Lock()
	fibo = &models.Fibonacci{
		Input:    x,
		Result:   fibResult,
		Duration: handler,
	}
	all = append(all, fibo)
	completed++
	mutex.Unlock()
}

func callFibonacci(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return callFibonacci(x-1) + callFibonacci(x-2)
}

func checkFibonacci(input uint64) {
	ok = 0
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

// Função que retorna o resultado do Fibonacci do número inserido
func GetNumber(c *fiber.Ctx) error {
	body := c.Params("input")
	input, err := strconv.ParseUint(body, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"done": false,
		})
	}
	checkFibonacci(input)

	count++
	if ok == 1 {

		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"done": true,
			"fib": fiber.Map{
				"": check,
			},
		})
	} else {
		go fibonacciCaller(input)
		time.Sleep(5 * time.Millisecond)
		if count != completed {
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"done": false,
			})
		} else {
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"done": true,
				"fib": fiber.Map{
					"": fibo,
				},
			})
		}
	}
}
