package main

import (
	"fmt"
	"reflect"
)

func _putchar(value interface{}) {
	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.String:
		fmt.Printf("%s", value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// Handle integer types
		if reflect.TypeOf(value) == reflect.TypeOf(rune(0)) {
			fmt.Printf("%c", value)
		} else {
			fmt.Printf("%d", value)
		}
	case reflect.Float32, reflect.Float64:
		// Handle float types
		fmt.Println(fmt.Sprintf("%f", value))
	case reflect.Bool:
		// Handle boolean type
		fmt.Println(fmt.Sprintf("%t", value))
	default:
		fmt.Println("Unsupported data type")
	}
}

func main() {
	fmt.Println("This is task 1 of the project")
	putchar()
	fmt.Println("This is task 2 of the project")
	print_alphabet()
	fmt.Println("This is task 3 of the project")
	X10_printalpha()
	fmt.Println("This is task 4 of the project")
	r := islower('H')
	_putchar(r)
	r = islower('o')
	_putchar(r)
	r = islower(108)
	_putchar(r)
	_putchar('\n')
	fmt.Println("This is task 5 of the project")
	r = isalpha('H')
	_putchar(r)
	r = isalpha('o')
	_putchar(r)
	r = isalpha(';')
	_putchar(r)
	r = isalpha(108)
	_putchar(r)
	_putchar('\n')
	fmt.Println("This is task 6 of the project")
	r = printSign(98)
	_putchar(',')
	_putchar(' ')
	_putchar(r)
	_putchar('\n')
	r = printSign(0)
	_putchar(',')
	_putchar(' ')
	_putchar(r)
	_putchar('\n')
	r = printSign(0xff)
	_putchar(',')
	_putchar(' ')
	_putchar(r)
	_putchar('\n')
	r = printSign(-2)
	_putchar(',')
	_putchar(' ')
	_putchar(r)
	_putchar('\n')
	fmt.Println("This is task 7 of the project")
	r = _abs(-1)
	fmt.Printf("%d\n", r)
	r = _abs(0)
	fmt.Printf("%d\n", r)
	r = _abs(1)
	fmt.Printf("%d\n", r)
	r = _abs(-98)
	fmt.Printf("%d\n", r)
	fmt.Println("This is task 8 of the project")
	printLast(98)
	printLast(0)
	r = printLast(-1024)
	_putchar(r)
	_putchar('\n')
	fmt.Println("This is task 9 of the project")
	jack_banuer()
	fmt.Println("This is task 10 of the project")
	timesTables()
	fmt.Println("This is task 11 of the project")
	n := _add(89, 9)
	fmt.Printf("%d\n", n)
	n = _add(102, 400)
	fmt.Printf("%d\n", n)
	fmt.Println("This is task 12 of the project")
	printTo98(0)
	printTo98(98)
	printTo98(111)
	printTo98(81)
	printTo98(-10)
}
