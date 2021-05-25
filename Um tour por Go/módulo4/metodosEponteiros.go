/* Comparando os dois programas anteriores, você pode notar que funções com um argumento de ponteiro deve ter um ponteiro:

var v Vertex
ScaleFunc(v, 5)  // Erro de compilação!
ScaleFunc(&v, 5) // OK

enquanto que os métodos com receptores ponteiro assumem um valor ou um ponteiro como o receptor quando eles são chamados:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK

Para a instrução v.Scale(5) ,
embora v seja um valor e não um ponteiro, o método com o receptor ponteiro é chamado automaticamente.
Ou seja, por conveniência, Go interpreta a declaração v.Scale(5) como (&v).Scale(5) uma vez que o método Scale tem um receptor de ponteiro.
*/

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
