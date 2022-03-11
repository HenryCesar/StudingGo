package controllers

import (
	"api-rest/models"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	all       = []*models.Fibonacci{}
	check     *models.Fibonacci
	fibo      *models.Fibonacci
	ok        int
	fibResult uint64
	mutex     sync.RWMutex
)

// Função que chama o Fibonacci
func fibonacciCaller(x uint64) {
	start := time.Now()

	fibResult = callFibonacci(x)

	elapsed := time.Since(start)
	duration := time.Duration(elapsed / time.Second)
	fmt.Println("elapsed do fibonacciCaller: ", elapsed, "duration do fibonacciCaller", duration)
	mutex.Lock()
	fibo = &models.Fibonacci{
		Input:    x,
		Result:   fibResult,
		Duration: duration,
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

func checkFibonacci(input uint64) int {
	ok = 0
	for _, c := range all {
		if c.Input == input {
			check = c
			ok = 1
			break
		}
	}
	return ok
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
	// Checa se o número inserido já existe
	checker := checkFibonacci(input)
	// Se sim, mostra pro usuário
	if checker == 1 {
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"done": true,
			"fib":  check,
		})
	} else { //Se não, roda a goroutine do fibonacci.
		go fibonacciCaller(input)
		time.Sleep(400 * time.Millisecond) //Espera 400ms pela resposta da goroutine
		checker2 := checkFibonacci(input)  //Checa se foi efetuado o fibonacci
		if checker2 == 1 {                 //Se sim, retorna o valor
			return c.Status(fiber.StatusFound).JSON(fiber.Map{
				"done": true,
				"fib":  check,
			})
		} else { // Se não, retorna "false" e continua rodando.
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"done": false,
			})
		}
	}
}
