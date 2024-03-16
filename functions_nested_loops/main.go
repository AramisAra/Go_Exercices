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
}
