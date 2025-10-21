package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"myFirstGoProject/pacote"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("--------------")
	fmt.Println(pacote.Bar)
	fmt.Println("--------------")

	pacote.PrintMinha()
	fmt.Println("--------------")

	digaOi()
	fmt.Println("--------------")

	fmt.Println(somar(1, 3))
	fmt.Println("--------------")

	a, b := swap(10, 20)
	fmt.Println(a, b)
	fmt.Println("--------------")

	res, rem := dividir(10, 3)
	fmt.Println(res, rem)
	fmt.Println("--------------")

	x := soma2(2)(1)
	fmt.Println("Tipo 1: ", x)
	q := soma2(2)
	w := q(1)
	fmt.Println("Tipo 2: ", w)
	fmt.Println("--------------")

	foo := func() {
		fmt.Println("Função Anonima!")
	}
	foo()
	fmt.Println("--------------")

	fmt.Println(soma3(10, 10, 10))
	fmt.Println("--------------")

	foo2("João Teste", 1, "Continua o Teste", 1, 2, 3)
	fmt.Println("--------------")

	variaveis()
	fmt.Println("--------------")

	takeInt32(10)
	takeInt64(10)

	fmt.Println("--------------")
	arrays()

	fmt.Println("--------------")
	loops()

	fmt.Println("--------------")
	ifconditional()

	fmt.Println("--------------")
	demoswitch()

	fmt.Println("--------------")
	fmt.Println(isWeekend(time.Now()))

	fmt.Println("--------------")
	fmt.Println(isWeekend2(time.Now()))

	fmt.Println("--------------")
	doDefer()

	fmt.Println("--------------")
	adGame()

}

func digaOi() {
	fmt.Println("Oi!")
}

func somar(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}
func soma2(a int) func(int) int {
	fmt.Println(a)
	return func(b int) int {
		fmt.Println(b)
		return a + b
	}
}

func soma3(nums ...int) int {
	var nossoOutPut int
	for _, n := range nums {
		nossoOutPut += n
	}
	return nossoOutPut
}

// Precisa ter ordem, como o inteiro nao se sabe a quantidade precisa ser declarado por ultimo
func foo2(a string, b byte, c string, n ...int) {
	// a = "João Teste"
	// b = 1
	// c = "Continua o Teste"
	// n = []int{1, 2, 3}
	fmt.Println(a, b, c, n)
}

var idade int

func variaveis() {
	var (
		nome      = "Pedro"
		sobrenome = "Silva"
	)
	fmt.Println(nome, sobrenome, idade)
}

//	boll
//
//	int	int8	int16	int32	int64
//	unint	uint8	uint16	uint32	uint64	uintptr
//
//	byte
//
//	rune
//
//	float32	float64
//
//	complex64 complex128
//
// string
//

func sum(a, b int) int {
	return a + b
}

func convert() {
	var x int = 65
	s := string(x)
	fmt.Println(s)
}

func takeInt32(x int32) {
	fmt.Println(x)
}
func takeInt64(x int64) {
	fmt.Println(x)
}

func arrays() {
	const n = 10
	arr := [n]int{5: 400}
	fmt.Println(arr)
}

func loops() {
	// var res int
	// for i := 0; i < 10; i++ {
	// 	res++
	// }
	// fmt.Println(res)

	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for indice, elemento := range arr {
	// 	fmt.Println(indice, elemento)
	// }

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, elemento := range arr {
		fmt.Println(elemento)
	}

	for i := range 10 {
		fmt.Println(i)
	}

	ponteiro := [3]int{1, 2, 3}
	for i, elem := range ponteiro {
		fmt.Println(&i, &elem)
	}

	const n = 10
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
func doError() error {
	return errors.New("error")
}

func ifconditional() {
	if 1 < 2 {
		fmt.Println("1 é menor que 2")
	}

	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1 {
		fmt.Println("Maior que zero!")
	} else {
		fmt.Println("Caiu no else")
	}

	boleano := true
	if boleano {
		fmt.Println("Aqui estou")
	}
	if err := doError(); err != nil {

	}
}
func do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Este é a opção padrão")
	}
}
func demoswitch() {
	do(1)
	do(2)
	do(3)
}

func isWeekend(dia time.Time) bool {
	switch {
	case dia.Weekday() > 0 && dia.Weekday() < 6:
		return false
	default:
		return true
	}
}

func isWeekend2(dia time.Time) bool {
	switch dia.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

func do2(x any) {
	switch t := x.(type) {
	case string:
		takeString(t)
	case int:
	case nil:
	}
}

func takeString(s string) {
	fmt.Println(s)
}

var arquivos = []string{"foo.txt", "bar.txt", "baz.txt"}

func doDefer() {
	defer fmt.Println("World")
	fmt.Println("Hello")

	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)
	x = 50
	fmt.Println(x)

	for _, f := range arquivos {
		func() {
			file, err := os.Open(f)
			if err != nil {
			}
			defer file.Close()
		}()
	}
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {

	}
	defer db.Close()
	return db, nil
}

func adGame() {
	fmt.Println("Jogo adivinhaçao, vamos jogar? ")
	fmt.Println("Um numero aleatorio sera sorteado. Tente acertar. O numeor é um inteiro entre 0 e 100:  ")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um numero inteiro")
			return
		}
		switch {
		case chuteInt < x:
			fmt.Println("Voce errou. O numero sorteado é maior que ", chuteInt)
		case chuteInt > x:
			fmt.Println("Voce errou. O numero sorteado é menor que ", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabens, voce acertou! 👏 "+
					"O numero era: %d\n"+
					"Voce acertou em %d tentativas. \n"+
					"Essas foram as suas tentativas: %v\n",
				x, i+1 , chutes[:i], )
			return

		}
		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Infelizmente, voce nao acertou o numero, que era: %d\n"+
			"Voce teve 10 tentativas. \n"+
			" Essas foram suas tentativas: %v\n", x, chutes)
}
