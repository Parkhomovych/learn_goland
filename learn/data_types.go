package learn

import (
	"fmt"
	"reflect"
	"unsafe"
)

// DataTypes демонструє основні типи даних у Go
func DataTypes() {
	fmt.Println("Основні типи даних у Go")
	fmt.Println("----------------------")

	// 1. Булеві типи
	fmt.Println("\n1. Булеві типи (bool):")
	boolTrue := true
	boolFalse := false
	fmt.Printf("true: %v, тип: %T\n", boolTrue, boolTrue)
	fmt.Printf("false: %v, тип: %T\n", boolFalse, boolFalse)

	// 2. Числові типи
	fmt.Println("\n2. Цілочисельні типи:")

	// Цілі числа зі знаком
	var intNum int = 42
	var int8Num int8 = 127
	var int16Num int16 = 32767
	var int32Num int32 = 2147483647
	var int64Num int64 = 9223372036854775807

	fmt.Printf("int: %v, тип: %T, розмір: %d байт\n", intNum, intNum, unsafe.Sizeof(intNum))
	fmt.Printf("int8: %v, тип: %T, розмір: %d байт\n", int8Num, int8Num, unsafe.Sizeof(int8Num))
	fmt.Printf("int16: %v, тип: %T, розмір: %d байт\n", int16Num, int16Num, unsafe.Sizeof(int16Num))
	fmt.Printf("int32: %v, тип: %T, розмір: %d байт\n", int32Num, int32Num, unsafe.Sizeof(int32Num))
	fmt.Printf("int64: %v, тип: %T, розмір: %d байт\n", int64Num, int64Num, unsafe.Sizeof(int64Num))

	// Цілі числа без знаку
	var uintNum uint = 42
	var uint8Num uint8 = 255
	var uint16Num uint16 = 65535
	var uint32Num uint32 = 4294967295
	var uint64Num uint64 = 18446744073709551615

	fmt.Printf("\nЦілі числа без знаку:\n")
	fmt.Printf("uint: %v, тип: %T, розмір: %d байт\n", uintNum, uintNum, unsafe.Sizeof(uintNum))
	fmt.Printf("uint8: %v, тип: %T, розмір: %d байт\n", uint8Num, uint8Num, unsafe.Sizeof(uint8Num))
	fmt.Printf("uint16: %v, тип: %T, розмір: %d байт\n", uint16Num, uint16Num, unsafe.Sizeof(uint16Num))
	fmt.Printf("uint32: %v, тип: %T, розмір: %d байт\n", uint32Num, uint32Num, unsafe.Sizeof(uint32Num))
	fmt.Printf("uint64: %v, тип: %T, розмір: %d байт\n", uint64Num, uint64Num, unsafe.Sizeof(uint64Num))

	// Спеціальні цілочисельні типи
	var byteNum byte = 255                                    // аліас для uint8
	var runeNum rune = 'Я'                                    // аліас для int32, представляє Unicode код символу
	var uintptrNum uintptr = uintptr(unsafe.Pointer(&intNum)) // для зберігання вказівника

	fmt.Printf("\nСпеціальні цілочисельні типи:\n")
	fmt.Printf("byte: %v, тип: %T, розмір: %d байт\n", byteNum, byteNum, unsafe.Sizeof(byteNum))
	fmt.Printf("rune: %v ('%c'), тип: %T, розмір: %d байт\n", runeNum, runeNum, runeNum, unsafe.Sizeof(runeNum))
	fmt.Printf("uintptr: %v, тип: %T, розмір: %d байт\n", uintptrNum, uintptrNum, unsafe.Sizeof(uintptrNum))

	// Числа з плаваючою точкою
	var float32Num float32 = 3.14159
	var float64Num float64 = 3.14159265358979323846

	fmt.Printf("\nЧисла з плаваючою точкою:\n")
	fmt.Printf("float32: %v, тип: %T, розмір: %d байт\n", float32Num, float32Num, unsafe.Sizeof(float32Num))
	fmt.Printf("float64: %v, тип: %T, розмір: %d байт\n", float64Num, float64Num, unsafe.Sizeof(float64Num))

	// Комплексні числа
	var complex64Num complex64 = 1 + 2i
	var complex128Num complex128 = 1.1 + 2.2i

	fmt.Printf("\nКомплексні числа:\n")
	fmt.Printf("complex64: %v, тип: %T, розмір: %d байт\n", complex64Num, complex64Num, unsafe.Sizeof(complex64Num))
	fmt.Printf("complex128: %v, тип: %T, розмір: %d байт\n", complex128Num, complex128Num, unsafe.Sizeof(complex128Num))

	// 3. Рядки
	fmt.Println("\n3. Рядки (string):")
	str := "Привіт, Go!"
	fmt.Printf("Рядок: %v, тип: %T, довжина: %d байт\n", str, str, len(str))
	fmt.Printf("Кількість символів (рун): %d\n", len([]rune(str)))

	// 4. Масиви
	fmt.Println("\n4. Масиви:")
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Масив: %v, тип: %T, довжина: %d\n", arr, arr, len(arr))

	// 5. Слайси (зрізи)
	fmt.Println("\n5. Слайси (зрізи):")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Слайс: %v, тип: %T, довжина: %d, ємність: %d\n", slice, slice, len(slice), cap(slice))

	// 6. Мапи (асоціативні масиви)
	fmt.Println("\n6. Мапи (асоціативні масиви):")
	m := map[string]int{
		"один": 1,
		"два":  2,
		"три":  3,
	}
	fmt.Printf("Мапа: %v, тип: %T, кількість елементів: %d\n", m, m, len(m))

	// 7. Структури
	fmt.Println("\n7. Структури:")
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Іван", Age: 30}
	fmt.Printf("Структура: %v, тип: %T\n", p, p)
	fmt.Printf("Поля структури: Name=%v, Age=%v\n", p.Name, p.Age)

	// 8. Вказівники
	fmt.Println("\n8. Вказівники:")
	x := 42
	ptr := &x
	fmt.Printf("Значення: %v, тип: %T\n", x, x)
	fmt.Printf("Вказівник: %v, тип: %T, значення за вказівником: %v\n", ptr, ptr, *ptr)

	// 9. Функції як тип
	fmt.Println("\n9. Функції як тип:")
	f := func(a, b int) int { return a + b }
	fmt.Printf("Функція: %v, тип: %T\n", f, f)
	fmt.Printf("Результат виклику функції f(5, 3): %v\n", f(5, 3))

	// 10. Інтерфейси
	fmt.Println("\n10. Інтерфейси:")
	var i interface{} = "Привіт"
	fmt.Printf("Інтерфейс: %v, тип: %T, тип значення: %v\n", i, i, reflect.TypeOf(i))

	i = 42
	fmt.Printf("Інтерфейс: %v, тип: %T, тип значення: %v\n", i, i, reflect.TypeOf(i))

	// 11. Канали
	fmt.Println("\n11. Канали:")
	ch := make(chan int)
	fmt.Printf("Канал: %v, тип: %T\n", ch, ch)

	// Демонстрація роботи з каналом у горутині
	go func() {
		ch <- 42 // відправка значення в канал
	}()

	value := <-ch // отримання значення з каналу
	fmt.Printf("Значення, отримане з каналу: %v\n", value)
}
