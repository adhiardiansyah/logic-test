package main

import (
	"bufio"
	"fmt"
	"math"
	"unicode"

	"os"
	"strings"
	"unicode/utf8"
)

func Reverse(s string) string {
	totalLength := len(s)
	buffer := make([]byte, totalLength)
	for i := 0; i < totalLength; {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		utf8.EncodeRune(buffer[totalLength-i:], r)
	}
	return string(buffer)
}

func isPalindrome(input string, caseSensitive bool) bool {
	if !caseSensitive {
		input = strings.ToLower(input)
	}

	// to deal with phrases, we need to eliminate the
	// spaces

	input = strings.TrimSpace(input)

	reverse := Reverse(input)

	if input == reverse {
		return true
	} else {
		return false
	}
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func reverseWords() {
	value := "I am A Great human"
	words := strings.Fields(value)
	result := ""
	for i, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			if unicode.IsUpper(rune(word[len(word)-1-i])) {
				result += strings.ToUpper(string(word[i]))
			} else {
				result += strings.ToLower(string(word[i]))
			}
		}

		if i != len(words)-1 {
			result += " "
		}
	}
	fmt.Println(result)
}

func NearestFibonaci() int {
	numbers := [...]int{15, 1, 3}
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	first, second := 0, 1

	third := first + second

	for third <= sum {
		first = second
		second = third
		third = first + second
	}

	if math.Abs(float64(third)-float64(sum)) >= math.Abs(float64(second)-float64(sum)) {
		return sum - second
	}

	return third - sum
}

func fizzbuzz() {
	var fizzbuzz int

	fmt.Print("Masukkan jumlah angka : ")
	fmt.Scan(&fizzbuzz)
	for i := 1; i <= fizzbuzz; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("CHECK PALINDROME")
	fmt.Print("Masukkan kalimat : ")

	consoleReader := bufio.NewReader(os.Stdin)
	answer, _ := consoleReader.ReadString('\n')

	// get rid of the extra newline character from ReadString() function
	answer = strings.TrimSuffix(answer, "\n")

	if isPalindrome(answer, false) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()

	fmt.Println("LEAP YEAR")

	var tahunawal int
	var tahunakhir int

	fmt.Print("Masukkan tahun awal : ")
	fmt.Scan(&tahunawal)
	fmt.Print("Masukkan tahun akhir : ")
	fmt.Scan(&tahunakhir)

	for i := tahunawal; i <= tahunakhir; i++ {
		if isLeap(i) {
			fmt.Print(i, ", ")
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()

	fmt.Println("REVERSE WORDS")
	reverseWords()

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()

	fmt.Println("NEAREST FIBONACI")
	fmt.Println(NearestFibonaci())

	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println()

	fmt.Println("FIZZBUZZ")
	fizzbuzz()
}
