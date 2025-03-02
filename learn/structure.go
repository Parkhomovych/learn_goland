package learn

import "fmt"

// Structure демонструє основні концепції структур у Go:
// - оголошення структур
// - вкладені структури
// - створення об'єктів структур
// - методи структур (з отримувачем-значенням та отримувачем-вказівником)
func Structure() {
	// Створюємо об'єкт Address за допомогою функції-конструктора
	address := MakeAddress("123 Main St", "Anytown", "CA", "12345")

	// Створюємо об'єкт Person з ім'ям, віком та адресою
	vlad := MakePerson("Vlad", 20, address)

	// Викликаємо методи структури Person
	vlad.SayHello() // Метод з отримувачем-значенням (копія)
	vlad.SayAge()   // Метод з отримувачем-вказівником (посилання)

	// Виводимо всю інформацію про об'єкт vlad
	fmt.Println("vlad", vlad)
}

// Person - структура, що представляє особу з ім'ям, віком та адресою
// Address вбудована як анонімне поле (вкладена структура)
type Person struct {
	Name    string
	Age     int
	Address // Вбудована структура (композиція замість наслідування)
}

// Address - структура, що представляє адресу з вулицею, містом, штатом та індексом
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// MakePerson - функція-конструктор для створення об'єкта Person
func MakePerson(name string, age int, address Address) Person {
	return Person{Name: name, Age: age, Address: address}
}

// MakeAddress - функція-конструктор для створення об'єкта Address
func MakeAddress(street string, city string, state string, zip string) Address {
	return Address{Street: street, City: city, State: state, Zip: zip}
}

// SayHello - метод структури Person з отримувачем-значенням (працює з копією)
func (p Person) SayHello() {
	fmt.Println("Hi, my name is ", p.Name)
}

// SayAge - метод структури Person з отримувачем-вказівником (працює з оригіналом)
func (p *Person) SayAge() {
	fmt.Println("I am ", p.Age, " years old")
}
