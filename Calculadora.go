package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entrada string, operador string) int {
	valores := strings.Split(entrada, operador)
	switch operador {
	case "+":
		return parsear(valores[0]) + parsear(valores[1])
	case "-":
		return parsear(valores[0]) - parsear(valores[1])
	case "*":
		return parsear(valores[0]) * parsear(valores[1])
	case "/":
		return parsear(valores[0]) / parsear(valores[1])
	default:
		return -1
	}
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	resultado := c.operate(entrada, operador)
	if resultado == -1 {
		fmt.Println("Error en la operación, operador no válido")
	} else {
		fmt.Println(resultado)
	}
}
