/* Para acessar o campo X de uma "struct" quando tivermos o ponteiro [p] da struct podemos escrever
[(*p).X]. No entanto, essa notação é incômoda, então a lingaugem nos permite escrever apenas [p.X],
sem a referência explícita. */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} //Denominando v um tipo Vertex, com os valores 1 e 2 para X e Y, respectivamente.
	p := &v           //O ponteiro [p] está recebendo como valor subjacente a variável [v] do tipo Vertex
	p.X = 1e9         //O ponteiro [p] possuí dentro dele a struct [Vertex], sendo assim possível acessar suas varíaveis.
	fmt.Println(v)    //As modificações feitas a partir do ponteiro, reescrevem os valores citados. Retornando ao processo de "indirecionamento".
}

// Campos de estruturas podem ser acessados através de um ponteiro de estrutura.
