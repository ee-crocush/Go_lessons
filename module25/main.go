package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Введите данные: ") // Исправленный текст
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("Вы ввели следующие данные: %s\n", input) // Исправленный текст
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
