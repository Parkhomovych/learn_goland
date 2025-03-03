package learn

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

// Errors демонструє різні способи роботи з помилками в Go
func Errors() {
	// 1. Базова обробка помилок
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Помилка при діленні:", err)
		return
	}
	fmt.Println("Результат ділення:", result)

	// Спроба ділення на нуль
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Помилка при діленні:", err)
		// Продовжуємо виконання, не повертаємось
	}

	// 2. Створення власних помилок
	// 2.1 За допомогою errors.New
	err = validateAge(-5)
	if err != nil {
		fmt.Println("Помилка валідації віку:", err)
	}

	// 2.2 За допомогою fmt.Errorf
	username := ""
	err = validateUsername(username)
	if err != nil {
		fmt.Println("Помилка валідації імені користувача:", err)
	}

	// 3. Робота з файлами та обробка помилок
	content, err := readFile("non_existent_file.txt")
	if err != nil {
		// 3.1 Перевірка типу помилки
		if os.IsNotExist(err) {
			fmt.Println("Файл не існує:", err)
		} else {
			fmt.Println("Інша помилка при читанні файлу:", err)
		}
	} else {
		fmt.Println("Вміст файлу:", content)
	}

	// 4. Власні типи помилок
	err = processTransaction(1001, 500)
	if err != nil {
		// 4.1 Перевірка типу власної помилки за допомогою type assertion
		if transErr, ok := err.(*TransactionError); ok {
			fmt.Printf("Помилка транзакції #%d: %s (код: %d)\n",
				transErr.TransactionID, transErr.Message, transErr.ErrorCode)
		} else {
			fmt.Println("Інша помилка:", err)
		}
	}

	// 5. Обгортання помилок (wrapping errors) - доступно з Go 1.13+
	err = executeBusinessLogic()
	if err != nil {
		// 5.1 Перевірка, чи містить помилка певний тип
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Помилка пов'язана з відсутністю файлу:", err)
		}

		// 5.2 Отримання конкретного типу помилки з ланцюжка
		var transErr *TransactionError
		if errors.As(err, &transErr) {
			fmt.Printf("Знайдено помилку транзакції в ланцюжку: %s (код: %d)\n",
				transErr.Message, transErr.ErrorCode)
		}

		// Виведення повного ланцюжка помилок
		fmt.Println("Повний ланцюжок помилок:", err)
	}
}

// 1. Базова функція, яка повертає помилку
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ділення на нуль неможливе")
	}
	return a / b, nil
}

// 2.1 Створення простої помилки за допомогою errors.New
func validateAge(age int) error {
	if age < 0 {
		return errors.New("вік не може бути від'ємним")
	}
	if age > 150 {
		return errors.New("вік занадто великий")
	}
	return nil
}

// 2.2 Створення форматованої помилки за допомогою fmt.Errorf
func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("ім'я користувача не може бути порожнім")
	}
	if len(username) < 3 {
		return fmt.Errorf("ім'я користувача '%s' занадто коротке (мінімум 3 символи)", username)
	}
	return nil
}

// 3. Функція для роботи з файлами та обробки помилок
func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err // Повертаємо помилку вище
	}
	return string(data), nil
}

// 4. Власний тип помилки
type TransactionError struct {
	TransactionID int
	ErrorCode     int
	Message       string
}

// Реалізація інтерфейсу error для власного типу помилки
func (e *TransactionError) Error() string {
	return fmt.Sprintf("транзакція #%d: %s (код помилки: %d)",
		e.TransactionID, e.Message, e.ErrorCode)
}

// Функція, яка повертає власний тип помилки
func processTransaction(id int, amount float64) error {
	if amount > 1000 {
		return &TransactionError{
			TransactionID: id,
			ErrorCode:     1001,
			Message:       "сума транзакції перевищує ліміт",
		}
	}
	// Успішна обробка
	return nil
}

// 5. Функція, що демонструє обгортання помилок (error wrapping)
func executeBusinessLogic() error {
	// Спроба виконати операцію
	err := performDatabaseOperation()
	if err != nil {
		// Обгортаємо помилку з додатковим контекстом
		return fmt.Errorf("помилка бізнес-логіки: %w", err)
	}
	return nil
}

func performDatabaseOperation() error {
	// Спроба прочитати конфігурацію
	err := readConfigFile()
	if err != nil {
		// Обгортаємо помилку з додатковим контекстом
		return fmt.Errorf("помилка операції з базою даних: %w", err)
	}
	return nil
}

func readConfigFile() error {
	// Симулюємо помилку відсутності файлу
	err := os.ErrNotExist
	if err != nil {
		// Обгортаємо стандартну помилку
		return fmt.Errorf("не вдалося прочитати конфігурацію: %w", err)
	}
	return nil
}

// 6. Додаткові приклади

// 6.1 Функція, що обробляє помилки при конвертації
func parseUserID(input string) (int, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("неправильний формат ID користувача: %w", err)
	}
	if id < 1 {
		return 0, errors.New("ID користувача повинен бути додатним числом")
	}
	return id, nil
}
