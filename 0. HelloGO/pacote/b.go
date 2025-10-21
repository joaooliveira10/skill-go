package pacote

import (
	"fmt"
	"myFirstGoProject/pacote/internal/foo"
)

var Bar string = "Hello bar"

func PrintMinha() {
	fmt.Println((foo.Minha))
}
