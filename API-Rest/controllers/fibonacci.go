package controllers

import (
	"api-rest/models"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	// Slice dos resultados do Fibonacci
	all = []*models.Fibonacci{}
	// Variável auxíliar/handler de um fibonacci
	fibo  *models.Fibonacci
	mutex sync.RWMutex
)

// Função que chama o Fibonacci
func fibonacciCaller(x uint64, ch chan int, ch2 chan int) {
	// Checa se há o Fibonacci do número inserido
	checker := checkFibonacci(x)
	if checker == 1 { // Se sim, finaliza a go routine
		ch <- checker
		return
	}
	start := time.Now()

	fibResult := callFibonacci(x)

	elapsed := time.Since(start)
	duration := time.Duration(elapsed / time.Second)
	// fmt.Println("elapsed do fibonacciCaller: ", elapsed, "duration do fibonacciCaller", duration)
	mutex.Lock()
	fibo = &models.Fibonacci{
		Input:    x,
		Result:   fibResult,
		Duration: duration,
	}
	all = append(all, fibo)
	mutex.Unlock()
	ch2 <- 1
}

func callFibonacci(x uint64) uint64 {
	if x <= 1 {
		return x
	}
	return callFibonacci(x-1) + callFibonacci(x-2)
}

func checkFibonacci(input uint64) int {
	var ok int
	ok = 0
	for _, c := range all {
		if c.Input == input {
			fibo = c
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
	chAux := make(chan int)
	chFib := make(chan int)
	go fibonacciCaller(input, chAux, chFib)
	for {
		select {
		case <-chAux: // Caso o canal auxiliar obtiver algo, retorna o resultado como "encontrado" junto ao conteúdo
			return c.Status(fiber.StatusFound).JSON(fiber.Map{
				"done": true,
				"fib":  fibo,
			})
		case <-time.After(time.Millisecond * 500): // Caso o tempo passe de 500 milisegundos, retorna falso.
			return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"done": false,
			})
		case <-chFib: // Caso o channel do Fibonacci obtiver algo, significa que houve o fibonacci do número, com isso retorna o resultado.
			return c.Status(fiber.StatusCreated).JSON(fiber.Map{
				"done": true,
				"fib":  fibo,
			})
		}
	}
}
