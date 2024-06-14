package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Gerador de Senhas Aleatórias")
	fmt.Println("============================")

	var length int
	fmt.Print("Digite o tamanho da senha desejada: ")
	fmt.Scan(&length)

	var useLowercase, useUppercase, useNumbers, useSymbols bool
	fmt.Print("Incluir letras minúsculas? (true/false): ")
	fmt.Scan(&useLowercase)
	fmt.Print("Incluir letras maiúsculas? (true/false): ")
	fmt.Scan(&useUppercase)
	fmt.Print("Incluir números? (true/false): ")
	fmt.Scan(&useNumbers)
	fmt.Print("Incluir símbolos? (true/false): ")
	fmt.Scan(&useSymbols)

	if !useLowercase && !useUppercase && !useNumbers && !useSymbols {
		fmt.Println("Erro: Pelo menos um tipo de caracter deve ser selecionado.")
		return
	}

	password := generatePassword(length, useLowercase, useUppercase, useNumbers, useSymbols)
	fmt.Println("Senha gerada:", password)
}

func generatePassword(length int, useLowercase, useUppercase, useNumbers, useSymbols bool) string {
	var chars []rune
	if useLowercase {
		chars = append(chars, generateCharset('a', 'z')...)
	}
	if useUppercase {
		chars = append(chars, generateCharset('A', 'Z')...)
	}
	if useNumbers {
		chars = append(chars, generateCharset('0', '9')...)
	}
	if useSymbols {
		symbols := []rune{'!', '@', '#', '$', '%', '^', '&', '*'}
		chars = append(chars, symbols...)
	}
	
	password := make([]rune, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func generateCharset(start, end rune) []rune {
	charset := make([]rune, end-start+1)
	for i := start; i <= end; i++ {
		charset[i-start] = i
	}
	return charset
}
