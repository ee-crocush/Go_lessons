package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// calculate выполняет математическую операцию.
func calculate(num1, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return num1 / num2, nil
	case "^":
		result := 1
		for i := 0; i < num2; i++ {
			result *= num1
		}
		return result, nil
	default:
		return 0, fmt.Errorf("неподдерживаемый оператор: %s", operator)
	}
}

// parseExpression извлекает числа и оператор из строки.
func parseExpression(line string) (int, string, int, error) {
	re := regexp.MustCompile(`(\d+)\s*([+\-*/^])\s*(\d+)\s*=?`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 4 {
		return 0, "", 0, fmt.Errorf("неверный формат выражения: %s", line)
	}

	num1, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования первого числа: %w", err)
	}

	operator := matches[2]

	num2, err := strconv.Atoi(matches[3])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования второго числа: %w", err)
	}

	return num1, operator, num2, nil
}

// processFile читает входной файл, вычисляет результаты и записывает их в выходной файл.
func processFile(inputFileName, outputFileName string) error {
	// Открываем входной файл
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return fmt.Errorf("ошибка открытия входного файла: %w", err)
	}
	defer inputFile.Close()

	// Создаем или очищаем выходной файл
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("ошибка создания/очистки выходного файла: %w", err)
	}
	defer outputFile.Close()

	// Создаем буферизированный writer
	bufferedWriter := bufio.NewWriter(outputFile)
	defer bufferedWriter.Flush()

	// Читаем входной файл построчно
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		// Парсим выражение
		num1, operator, num2, err := parseExpression(line)
		if err != nil {
			log.Printf("ошибка парсинга строки '%s': %v", line, err)
			continue // Переходим к следующей строке
		}

		// Вычисляем результат
		result, err := calculate(num1, num2, operator)
		if err != nil {
			log.Printf("ошибка вычисления выражения '%d %s %d': %v", num1, operator, num2, err)
			continue // Переходим к следующей строке
		}

		// Записываем результат в выходной файл
		resultString := fmt.Sprintf("%d%s%d=%d\n", num1, operator, num2, result)
		_, err = bufferedWriter.WriteString(resultString)
		if err != nil {
			return fmt.Errorf("ошибка записи в выходной файл: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения входного файла: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Использование: program <входной файл> <выходной файл>")
		return
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	err := processFile(inputFileName, outputFileName)
	if err != nil {
		log.Fatalf("Ошибка обработки файла: %v", err)
	}

	fmt.Println("Обработка завершена. Результаты записаны в", outputFileName)
}
