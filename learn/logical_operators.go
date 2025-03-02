package learn

import "fmt"

// LogicalOperators демонструє роботу логічних операторів у Go
func LogicalOperators() {
	fmt.Println("Логічні оператори в Go")
	fmt.Println("---------------------")

	// Основні логічні оператори в Go:
	// && (логічне І, AND)
	// || (логічне АБО, OR)
	// !  (логічне НЕ, NOT)

	// Логічний оператор && (AND)
	fmt.Println("\n1. Логічний оператор && (AND):")
	fmt.Printf("true && true = %v\n", true && true)     // true
	fmt.Printf("true && false = %v\n", true && false)   // false
	fmt.Printf("false && true = %v\n", false && true)   // false
	fmt.Printf("false && false = %v\n", false && false) // false

	// Логічний оператор || (OR)
	fmt.Println("\n2. Логічний оператор || (OR):")
	fmt.Printf("true || true = %v\n", true || true)     // true
	fmt.Printf("true || false = %v\n", true || false)   // true
	fmt.Printf("false || true = %v\n", false || true)   // true
	fmt.Printf("false || false = %v\n", false || false) // false

	// Логічний оператор ! (NOT)
	fmt.Println("\n3. Логічний оператор ! (NOT):")
	fmt.Printf("!true = %v\n", !true)   // false
	fmt.Printf("!false = %v\n", !false) // true

	// Комбінування логічних операторів
	fmt.Println("\n4. Комбінування логічних операторів:")
	fmt.Printf("true && !false = %v\n", true && !false)                                         // true
	fmt.Printf("!true || !false = %v\n", !true || !false)                                       // true
	fmt.Printf("!(true && false) = %v\n", !(true && false))                                     // true
	fmt.Printf("(true || false) && (true && false) = %v\n", (true || false) && (true && false)) // false

	// Короткозамкнена оцінка (short-circuit evaluation)
	fmt.Println("\n5. Короткозамкнена оцінка (short-circuit evaluation):")

	// Для оператора &&: якщо перший операнд false, другий не обчислюється
	x := false
	y := true

	// Функція, яка повертає true і виводить повідомлення
	checkAndPrint := func(name string) bool {
		fmt.Printf("Функція checkAndPrint викликана з аргументом %s\n", name)
		return true
	}

	fmt.Println("\nПриклад з &&:")
	result1 := x && checkAndPrint("test1") // checkAndPrint не викликається через короткозамкнену оцінку
	fmt.Printf("Результат x && checkAndPrint(\"test1\"): %v\n", result1)

	result2 := y && checkAndPrint("test2") // checkAndPrint викликається
	fmt.Printf("Результат y && checkAndPrint(\"test2\"): %v\n", result2)

	// Для оператора ||: якщо перший операнд true, другий не обчислюється
	fmt.Println("\nПриклад з ||:")
	result3 := x || checkAndPrint("test3") // checkAndPrint викликається
	fmt.Printf("Результат x || checkAndPrint(\"test3\"): %v\n", result3)

	result4 := y || checkAndPrint("test4") // checkAndPrint не викликається через короткозамкнену оцінку
	fmt.Printf("Результат y || checkAndPrint(\"test4\"): %v\n", result4)

	// Практичне використання короткозамкненої оцінки
	fmt.Println("\n6. Практичне використання короткозамкненої оцінки:")

	// Перевірка на nil перед доступом до поля або методу
	type Person struct {
		Name string
	}

	var p *Person // nil вказівник

	// Безпечна перевірка: спочатку перевіряємо на nil, потім доступаємося до поля
	if p != nil && p.Name == "Іван" {
		fmt.Println("Це Іван")
	} else {
		fmt.Println("Це не Іван або p є nil")
	}

	// Ініціалізуємо p
	p = &Person{Name: "Іван"}

	// Тепер перевірка пройде успішно
	if p != nil && p.Name == "Іван" {
		fmt.Println("Це Іван")
	} else {
		fmt.Println("Це не Іван або p є nil")
	}

	// Використання || для встановлення значення за замовчуванням
	var input string = "" // Припустимо, це введення користувача
	name := input
	if name == "" {
		name = "Гість" // Встановлюємо значення за замовчуванням
	}
	fmt.Printf("Привіт, %s!\n", name)

	// Те саме з використанням ||
	input = "" // Скидаємо введення
	// Якщо input порожній (false при порівнянні з непорожнім рядком), використовуємо "Гість"
	nameWithOr := "Гість"
	if input != "" {
		nameWithOr = input
	}
	fmt.Printf("Привіт, %s! (з використанням ||)\n", nameWithOr)
}
