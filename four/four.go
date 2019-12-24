package four

import (
	"fmt"
	"strconv"
	"strings"
)

func passwordLength(password int) bool {
	return len(strconv.Itoa(password)) == 6
}

func passwordIncreasing(password int) bool {
	var prev = -1
	for _, val := range strconv.Itoa(password) {
		cur, err := strconv.Atoi(string(val))
		if err != nil {
			return false
		}
		if cur < prev {
			return false
		}
		prev = cur
	}
	return true
}

func passwordRepeatDigits(password int) bool {
	var prev = 'c'
	for _, val := range strconv.Itoa(password) {
		if prev == val {
			return true
		}
		prev = val
	}
	return false
}

func passwordRepeatPairs(password int) bool {
	var prev = 'c' // Non Numeric Char, cause the first is never a match
	var repeatCount = 1
	for _, val := range strconv.Itoa(password) {
		if prev != val {
			if repeatCount == 2 {
				return true
			}
			repeatCount = 1
		} else {
			repeatCount += 1
		}
		prev = val
	}
	if repeatCount == 2 {
		return true
	}
	return false
}

func getBounds(input []string) (int, int) {
	split := strings.Split(input[0], "-")
	lower, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0
	}
	upper, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0
	}
	return lower, upper
}

func generatePasswordsOne(lower int, upper int) []int {
	var passwords []int
	for i := lower; i <= upper; i++ {
		if passwordIncreasing(i) && passwordLength(i) && passwordRepeatDigits(i) {
			passwords = append(passwords, i)
		}
	}
	return passwords
}

func generatePasswordsTwo(lower int, upper int) []int {
	var passwords []int
	for i := lower; i <= upper; i++ {
		if passwordIncreasing(i) && passwordLength(i) && passwordRepeatPairs(i) {
			passwords = append(passwords, i)
		}
	}
	return passwords
}

func Run(input []string) {
	lower, upper := getBounds(input)
	passwords := generatePasswordsOne(lower, upper)
	fmt.Printf("Passwords between %v and %v: %v\n", lower, upper, len(passwords))
	passwords = generatePasswordsTwo(lower, upper)
	fmt.Printf("Passwords w/ Pairs between %v and %v: %v\n", lower, upper, len(passwords))
}
