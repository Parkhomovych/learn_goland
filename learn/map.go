package learn

import "fmt"

// Map демонструє роботу з асоціативними масивами (map) в Go
// Map - це невпорядкована колекція пар ключ-значення
func Map() {
	// Створення map за допомогою функції make
	// make(map[KeyType]ValueType, [capacity])
	// 1. Виділяє пам'ять для хеш-таблиці
	// 2. Ініціалізує внутрішні структури даних
	// 3. Встановлює початкову ємність (якщо вказана)
	// 4. Повертає ініціалізовану Map, готову до використання

	// Приклад 1: Створення map без вказання ємності
	users := make(map[string]int)

	// Додавання елементів у map
	users["Alice"] = 25
	users["Bob"] = 30
	users["Charlie"] = 22

	// Отримання значення за ключем
	aliceAge := users["Alice"]
	fmt.Println("Alice's age:", aliceAge)

	// Перевірка наявності ключа
	age, exists := users["Dave"]
	if exists {
		fmt.Println("Dave's age:", age)
	} else {
		fmt.Println("Dave not found in the map")
	}

	// Видалення елемента з map
	delete(users, "Bob")

	// Перебір всіх елементів map
	fmt.Println("All users:")
	for name, age := range users {
		fmt.Printf("%s: %d years old\n", name, age)
	}

	// Приклад 2: Створення map із вказаною ємністю
	// Корисно, коли заздалегідь відома приблизна кількість елементів
	cities := make(map[string]string, 10)

	cities["Kyiv"] = "Ukraine"
	cities["Paris"] = "France"
	cities["London"] = "UK"
	cities["Rome"] = "Italy"

	fmt.Println("\nCities and their countries:")
	for city, country := range cities {
		fmt.Printf("%s is in %s\n", city, country)
	}

	// Приклад 3: Різниця між nil map і порожньою map

	// Nil map (не ініціалізована)
	var nilMap map[string]int
	// nilMap["test"] = 1 // Це викличе panic: assignment to entry in nil map

	// Перевірка, чи є map nil
	if nilMap == nil {
		fmt.Println("\nnilMap is nil and cannot be used without initialization")
	}

	// Порожня map (ініціалізована)
	emptyMap := make(map[string]int)
	emptyMap["test"] = 1 // Це працює нормально
	fmt.Println("Value in emptyMap:", emptyMap["test"])

	// Приклад 4: Використання map як набору (set)
	// Go не має вбудованого типу set, але map можна використовувати як set
	uniqueNames := make(map[string]bool)

	// Додавання елементів у "set"
	uniqueNames["Alice"] = true
	uniqueNames["Bob"] = true
	uniqueNames["Alice"] = true // Дублікат, але в map залишиться тільки один запис

	// Перевірка наявності елемента в "set"
	_, isAliceInSet := uniqueNames["Alice"]
	fmt.Println("\nIs Alice in the set?", isAliceInSet)

	// Виведення всіх унікальних імен
	fmt.Println("Unique names:")
	for name := range uniqueNames {
		fmt.Println(name)
	}
}
