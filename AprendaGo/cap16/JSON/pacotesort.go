package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"abóbora", "maça", "laranja", "berinjela"} //sort.Strings

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)

	si := []int{1, 26, 81, 45, 75, 698, 70, 781, 19}

	fmt.Println(si)

	sort.Ints(si)

	fmt.Println(si)
}
