package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	kafka_broker "kafka-cli/pkg/kafka"
	"kafka-cli/pkg/prompt"
	"os"
	"strings"
)

func main() {
	broker := kafka_broker.NewBroker()
	reader := bufio.NewReader(os.Stdin)

	const (
		showLabel    = "Показать топики"
		сreateLabel  = "Создать топик"
		sendMesLabel = "Отправить сообщение"
		exitLabel    = "Выход"
	)

	fmt.Printf("Kafka Producer запущен\nПодключен к: %s\n\n", broker.Address())

	for {
		menu := promptui.Select{
			Label: "Выберите команду",
			Items: []string{showLabel, сreateLabel, sendMesLabel, exitLabel},
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
			} else {
				fmt.Println("Доступные топики:")
				for _, t := range topics {
					fmt.Println(" -", t)
				}
			}
		case сreateLabel:
			fmt.Println("Введите название топика:")
			topic, _ := reader.ReadString('\n')
			topic = strings.TrimSpace(topic)
			err := broker.CreateTopic(topic)
			if err != nil {
				fmt.Println("Произошла ошибка при создании топика: ", err)
			} else {
				fmt.Println("Топик успешно создан")
			}
		case sendMesLabel:
			topics, err := broker.ListTopics()
			if err != nil {
				fmt.Println("Ошибка получения топиков: ", err)
				continue
			}

			if len(topics) == 0 {
				fmt.Println("Нет доступных топиков для записи")
				continue
			}

			topic, err := prompt.PromptSelect("Выберите топик", topics)
			if err != nil {
				fmt.Println("Ошибка выбора топика: ", err)
				continue
			}

			fmt.Printf("Введите сообщение для отправки в топик \"%s\"\n", topic)
			msg, _ := reader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			err = broker.SendMessage(topic, msg)
			if err != nil {
				fmt.Println("Ошибка при отправке сообщения: ", err)
			} else {
				fmt.Printf("Сообщение успешно отправлено в топик \"%s\"\n", topic)
			}
		case exitLabel:
			fmt.Println("До встречи!")
			return
		}
	}
}
