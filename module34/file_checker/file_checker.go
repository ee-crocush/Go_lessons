package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите путь к директории: ")
	dirPath, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	dirPath = strings.TrimSpace(dirPath)

	// Получаем список файлов и директорий
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	// Создаем или перезаписываем файл LS.txt
	outputFile, err := os.Create("LS.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, entry := range entries {
		// Получаем доп. информацию (в частности, размер файла)
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(writer, "%s (не удалось получить размер) [UNKNOWN]\n", entry.Name())
			continue
		}

		fileType := "FILE"
		if info.IsDir() {
			fileType = "DIRECTORY"
		}

		_, err = writer.WriteString(fmt.Sprintf("%s (%d bytes) [%s]\n", info.Name(), info.Size(), fileType))
		if err != nil {
			panic(err)
		}
	}

	if err := writer.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("Список файлов и папок успешно записан в LS.txt")
}
