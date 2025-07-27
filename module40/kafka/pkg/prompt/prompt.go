package prompt

import "github.com/manifoldco/promptui"

// PromptSelect представляет функцию для отображения выпадающего списка с выбором опции пользователя.// Эта функция оборачивает логику взаимодействия с пользователем через promptui,
// возвращает выбранный элемент из списка или ошибку, если выбор не удался.
func PromptSelect(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, results, err := prompt.Run()

	return results, err
}
