/* O case default em um select é executado se nenhum outro caso está pronto.

Utilize um case default para tentar enviar ou receber sem bloqueio:

	select{
	case i := <-c:
		// use i
	default:
		//recebendo c bloqueado
	} */

package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
