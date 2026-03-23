package main

import "fmt"

func main() {
	var productName string = "Iphone"
	var productPrice int = 10000
	fmt.Println(productName, productPrice)

	loops_demo()
	arrays_demo()
	maps_demo()
	demo_pointers()

	x, y := check_even_odd(10)

	fmt.Println(x, " and the return value is ", y)
}

func loops_demo() {
	for i := 0; i < 3; i++ {
		fmt.Println("Loop iteration", i)
	}

	for i := range 3 {
		fmt.Println("Range loop itteration: ", i)
	}

	for i, char := range "sid" {
		fmt.Println("String range loop iteration: ", i, char)
	}

	// Below is the while using for loop

	j := 10

	for j > 0 {
		if j == 3 {
			j--
			continue
		} else if j == 5 {
			fmt.Println("Skipping 5")
			j--
			continue
		} else {
			fmt.Println("while loop iteration: ", j)
			j--
			continue
		}
	}

	for {
		fmt.Println("Infinite loop iteration")
		break
	}

}

func arrays_demo() {

	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	var arr1 []int // slice of empty integers

	arr1 = append(arr1, 1, 2, 3, 4, 5)

	fmt.Println("Array 1:", arr1)

	// myarr := []int{1, 2, 3, 4, 5} // Declare and initialize a slice with values

	// fmt.Println(myarr)
}

func maps_demo() {
	productPrices := map[string]int{
		"iPhone 15 Pro": 1000,
		"iPhone 15":     800,
		"iPhone 14":     600,
	}

	fmt.Println("Product Prices: ", productPrices)

	customMap := map[string]string{}

	fmt.Println("Custom Map: ", customMap)

	emptyMap := make(map[string]int)

	emptyMap["key1"] = 100
	emptyMap["key2"] = 200
	emptyMap["key3"] = 300

	fmt.Println("empty map: ", emptyMap)
}

func check_even_odd(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 0
	} else {
		return "Odd", 1
	}
}

func demo_pointers() {
	i := 120

	var ptr *int = &i // this poitner can store the address of an integer variable

	ptr1 := &i

	fmt.Println("Value of i: ", i, " Pointer to i: ", ptr, "Pointer1 to i: ", ptr1)

	fmt.Println("Value present at pointer ", *ptr)
}
