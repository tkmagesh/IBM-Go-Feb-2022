package main

import (
	"fmt"
	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Yellow("Calculator Operations")
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println(calc.OpCount())

	color.Red("Utils Operations")
	fmt.Println(utils.IsEven(71))
	fmt.Println(utils.IsOdd(71))

}
