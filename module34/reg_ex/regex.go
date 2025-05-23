package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	minLength = 8
)

func main() {

	//reader := bufio.NewReader(os.Stdin)
	//checkPassword(reader)

	contentBytes, err := os.ReadFile("./index.html")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`class="([a-zA-Z0-9_\-\s]+)"`)

	submatches := re.FindAllStringSubmatch(string(contentBytes), -1)

	for _, s := range submatches {
		classes := strings.Split(s[1], " ")
		for _, c := range classes {
			fmt.Println("Найден класс", c)
		}
	}

}

func emailCheck(reader *bufio.Reader) {
	fmt.Print("Введите email: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// удаляем символ перевода на новую строку
	str = strings.TrimSpace(str)

	emailRegex := regexp.MustCompile(`^[a-zA-Z][\w.-]+@[\w.-]+$`)
	isMatch := emailRegex.MatchString(str)

	fmt.Printf("Ввод \"%s\", совпадение с шаблоном: %v", str, isMatch)
}

func checkPassword(reader *bufio.Reader) {
	fmt.Print(
		"Введите пароль " +
			"(должен содержать строчные и прописные буквы, " +
			"цифры, по крайней мере один спец.символ " +
			"и быть длинной не менее 8 символов): ",
	)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// удаляем символ перевода на новую строку
	str = strings.TrimSpace(str)

	// проверка длины пароля
	lenRegex := regexp.MustCompile(fmt.Sprintf(`^.{%d,}$`, minLength))
	if !lenRegex.MatchString(str) {
		fmt.Printf("Ошибка! Длина пароля менее %d\n", minLength)
		return
	}

	// проверка наличия обязательных символов
	if !regexp.MustCompile(`[a-z]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать строчные буквы")
		return
	}
	if !regexp.MustCompile(`[A-Z]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать прописные буквы")
		return
	}
	if !regexp.MustCompile(`[0-9]+`).MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать цифры")
		return
	}

	mostPopularPassword := []string{
		"Qq123456",
		"Qwerty123",
	}
	join := strings.Join(mostPopularPassword, "|")
	weakPassRegex := regexp.MustCompile(fmt.Sprintf("^(%s)$", join))
	if weakPassRegex.MatchString(str) {
		fmt.Println("Предупреждение! Очень слабый пароль, придумайте другой")
		return
	}

	specCharRegex := regexp.MustCompile(`[!@#$%^&*()\-=+,./\\]+`)
	if !specCharRegex.MatchString(str) {
		fmt.Println("Ошибка! Пароль должен содержать спец.символ")
		return
	}

	fmt.Println("Введен корректный пароль")
}
