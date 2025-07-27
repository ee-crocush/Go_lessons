package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	kafka_broker "kafka-cli/pkg/kafka"
	"kafka-cli/pkg/prompt"
)

func main() {
	broker := kafka_broker.NewBroker()

	const (
		showLabel = "Показать топики"
		readLabel = "Прочитать сообщения из топика"
		exitLabel = "Выход"
	)

	fmt.Printf("Kafka Consumer запущен\nПодключен к: %s\n\n", broker.Address())

	for {
		menu := promptui.Select{
			Label: "Выберите команду",
			Items: []string{showLabel, readLabel, exitLabel},
		}

		_, choice, err := menu.Run()
		if err != nil {
			fmt.Printf("Ошибка выбора команды: %v\n", err)
			continue
		}

		switch choice {
		case showLabel:
			topics, err := broker.ListTopics()
			if err != nil {
				fmt.Println("Ошибка чтения топиков: ", err)
				continue
			}

			if len(topics) == 0 {
				fmt.Println("Нет доступных топиков")
			}

			if len(topics) == 0 {
				fmt.Println("Нет доступных топиков")
			} else {
				fmt.Println("Доступные топики:")
				for _, t := range topics {
					fmt.Println(" -", t)
				}
			}
		case readLabel:
			topics, err := broker.ListTopics()
			if err != nil {
				fmt.Println("Ошибка получения топиков: ", err)
				continue
			}

			if len(topics) == 0 {
				fmt.Println("Нет доступных топиков для чтения\n")
				continue
			}

			topic, err := prompt.PromptSelect("Выберите топик", topics)
			if err != nil {
				fmt.Println("Ошибка выбора топика: ", err)
				continue
			}

			if err := broker.ReadMessage(topic); err != nil {
				fmt.Printf("Ошибка чтения сообщений: %v\n", err)
			}
			fmt.Println()
		case exitLabel:
			fmt.Println("До встречи!")
			return
		}
	}
}
