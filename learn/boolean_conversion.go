package learn

import "fmt"

// BooleanConversion демонструє особливості приведення до булевого типу в Go
func BooleanConversion() {
	fmt.Println("Приведення до булевого типу в Go")
	fmt.Println("--------------------------------")

	// В Go немає автоматичного приведення типів до булевого значення
	// Наприклад, наступний код не скомпілюється:
	// if "" { // Помилка: non-boolean condition in if statement
	// 	fmt.Println("Порожній рядок вважається true")
	// }

	// if nil { // Помилка: use of untyped nil in if statement
	// 	fmt.Println("nil вважається true")
	// }

	// if 0 { // Помилка: non-boolean condition in if statement
	// 	fmt.Println("0 вважається true")
	// }

	// Правильний спосіб перевірки порожнього рядка
	emptyString := ""
	if emptyString == "" {
		fmt.Println("Порожній рядок перевіряється через порівняння: emptyString == \"\"")
	}

	// Перевірка nil для інтерфейсів та вказівників
	var nilPointer *int
	if nilPointer == nil {
		fmt.Println("nil перевіряється через порівняння: nilPointer == nil")
	}

	// Перевірка для слайсів, мап та каналів
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nil слайс перевіряється через порівняння: nilSlice == nil")
	}

	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("nil мапа перевіряється через порівняння: nilMap == nil")
	}

	// Перевірка довжини слайсів та мап
	emptySlice := []int{}
	if len(emptySlice) == 0 {
		fmt.Println("Порожній слайс перевіряється через len(emptySlice) == 0")
	}

	// Перевірка булевих значень
	boolValue := false
	if !boolValue {
		fmt.Println("Булеве значення false перевіряється через !boolValue")
	}

	// Використання логічних операторів
	a, b := true, false
	if a && !b {
		fmt.Println("Логічні оператори: a && !b")
	}
}
