package main

// Mencari Bilangan Ganjil Genap dari 1 - 100
// func main() {
//     fmt.Println("Bilangan Genap:")
//     for i := 1; i <= 100; i++ {
//         if i%2 == 0 {
//             fmt.Print(i, " ")
//         }
//     }

//     fmt.Println("\n\nBilangan Ganjil:")
//     for i := 1; i <= 100; i++ {
//         if i%2 != 0 {
//             fmt.Print(i, " ")
//         }
//     }
// }

// Mencari Bilangan Prima
// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n <= 3 {
// 		return true
// 	}
// 	if n%2 == 0 || n%3 == 0 {
// 		return false
// 	}

// 	i := 5
// 	for i*i <= n {
// 		if n%i == 0 || n%(i+2) == 0 {
// 			return false
// 		}
// 		i += 6
// 	}

// 	return true
// }

// func main() {
// 	fmt.Println("Bilangan prima antara 1 dan 100:")

// 	for i := 1; i <= 100; i++ {
// 		if isPrime(i) {
// 			fmt.Printf("%d ", i)
// 		}
// 	}
// }

// Fibonacci
// func main() {
//     a := 0;
// 	b := 1;

//     fmt.Print("Deret Fibonacci: ")

//     for i := 0; i < 10; i++ {
//         fmt.Print(a, " ")

//         a, b = b, a+b
//     }

//     fmt.Println()
// }

// Reverse String
// func reverseString(input string) string {
//     runes := []rune(input)
//     n := len(runes)
//     for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }
//     return string(runes)
// }

// func main() {
//     inputString := "Hello, World!"
//     reversedString := reverseString(inputString)
//     fmt.Println("Original String:", inputString)
//     fmt.Println("Reversed String:", reversedString)
// }

// FizzBuzz
// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

//Bubble Sort
// func bubbleSort(arr []int) {
//     n := len(arr)
//     for i := 0; i < n-1; i++ {
//         for j := 0; j < n-i-1; j++ {
//             if arr[j] > arr[j+1] {
//                 arr[j], arr[j+1] = arr[j+1], arr[j]
//             }
//         }
//     }
// }

// func main() {
//     arr := []int{64, 34, 25, 12, 22, 11, 90}
//     fmt.Println("Array sebelum pengurutan:", arr)

//     bubbleSort(arr)

//     fmt.Println("Array setelah pengurutan:", arr)
// }


