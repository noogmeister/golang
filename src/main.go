package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//end pakage level variables.

func oneD6andTwoD8s() (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6), rand.Intn(8), rand.Intn(8)
}

func threeD6() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(6) + rand.Intn(6) + rand.Intn(6)
}
func textUserInterface() {
	reader := bufio.NewReader(os.Stdin)

	for true {

		fmt.Print("type 1 or 2: ")
		input, _ := reader.ReadString('\n')
		if strings.TrimRight(input, "\r\n") == "1" {
			fmt.Printf("the total of 3d6 is %d\n", threeD6()+3)
		}
		if strings.TrimRight(input, "\r\n") == "2" {
			a, b, c := oneD6andTwoD8s()
			fmt.Printf("\n1 d6 is %d \nand 2 d8s are %d, \nand %d\n", a+1, b+1, c+1)
		}
		if strings.TrimRight(input, "\r\n") == "q" {
			fmt.Print("\nexiting program")
			break
		}
	}
}
func main() {
	textUserInterface()
}
