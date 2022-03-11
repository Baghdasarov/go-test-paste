package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	str := "23-abw-48-caba-56-asd"
	fmt.Printf("Given string: %s\n", str)
	fmt.Printf("testValidity: %v\n", testValidity(str))
	fmt.Printf("averageNumber : %v\n", averageNumber(str))
	fmt.Printf("wholeStory: %v\n", wholeStory(str))
	shortest, longest, avg, result := storyStats(str)
	fmt.Printf("storyStats:"+
		"\n\tshortest: %v"+
		"\n\tlongest: %v"+
		"\n\tavg: %v"+
		"\n\tresult: %v\n", shortest, longest, avg, result)
	fmt.Printf("generateString:"+
		"\n\tflag true: %v"+
		"\n\tflag false: %v\n", generateString(true), generateString(false))
}

//
//  testValidity
//  @Description: That function takes the string value, and returns true if the given string
//				  is a sequence of numbers followed by dash followed by text,
//				  or false if the string does not comply with the format
//  @param str
//  @return bool
//  @estimated time: 30 minutes
//  @used time: 25 minutes
//
func testValidity(str string) bool {
	if len(str) < 3 {
		return false
	}

	items := strings.Split(str, "-")

	if _, err := getDigit(items[len(items)-1]); err == nil {
		return false
	}

	for i := 0; i < len(items); i++ {
		if items[i] == "" {
			return false
		}
		if _, currErr := getDigit(items[i]); currErr != nil {
			if i == 0 {
				return false
			}
			_, prevErr := getDigit(items[i-1])
			if i == len(items)-1 {
				if prevErr != nil {
					return false
				}
				return true
			}
			_, nextErr := getDigit(items[i+1])
			if prevErr != nil || nextErr != nil {
				return false
			}
		}
	}
	return true
}

//
//  averageNumber
//  @Description: That function takes the string, filters all numbers from there and returns the average
//  @param str
//  @return float64
//  @estimated time: 10 minutes
//  @used time: 7 minutes
//
func averageNumber(str string) float64 {
	items := strings.Split(str, "-")
	sum := 0.0
	numOfDigits := 0
	for i := 0; i < len(items); i++ {
		if num, err := getDigit(items[i]); err == nil {
			sum += float64(num)
			numOfDigits++
		}
	}
	return sum / float64(numOfDigits)
}

//
//  wholeStory
//  @Description: That function that takes the string, filters all strings/words from there and returns a text
// 				  that is composed of all those words separated by spaces
//  @param str
//  @return string
//  @estimated time: 10 minutes
//  @used time: 10 minutes
//
func wholeStory(str string) string {
	items := strings.Split(str, "-")
	var result string
	for i := 0; i < len(items); i++ {
		if _, err := getDigit(items[i]); err != nil {
			result += items[i]
			if i != len(items)-1 {
				result += " "
			}
		}
	}
	return result
}

//
//  storyStats
//  @Description: Function that find and return the shortest, longest words, the average of word length and the slice
//				  of all words from the text that have the length the same as the average length rounded up and down
//  @param str
//  @return shortest
//  @return longest
//  @return avg
//  @return result
//  @estimated time: 30 minutes
//  @used time: 30 minutes
//
func storyStats(str string) (shortest string, longest string, avg float64, result []string) {
	items := strings.Split(str, "-")
	stringsSlice := make([]string, 0)
	sum := 0.0

	for i := 0; i < len(items); i++ {
		if _, err := getDigit(items[i]); err != nil {
			stringsSlice = append(stringsSlice, items[i])
			sum += float64(len(items[i]))
		}
	}

	avg = sum / float64(len(stringsSlice))

	for _, word := range stringsSlice {
		if len(word) < len(shortest) || shortest == "" {
			shortest = word
		}
		if len(word) > len(longest) {
			longest = word
		}

		if float64(len(word)) == math.Floor(avg) || float64(len(word)) == math.Ceil(avg) {
			result = append(result, word)
		}
	}

	return shortest, longest, avg, result
}

//
//  generateString
//  @Description: Function which takes a boolean flag and generates random correct strings that match this pattern
//  			  (sequence of numbers followed by dash followed by text) if the parameter is true, and random invalid
// 				  strings if the flag is false
//  @param flag
//  @return result
//  @estimated time: 30 minutes
//  @used time: 25 minutes
//
func generateString(flag bool) (result string) {
	if flag {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				result += fmt.Sprintf("%v-", RandomInt(1, 50))
				continue
			}
			result += RandomString(RandomInt(1, 10))
			if i != 9 {
				result += "-"
			}
		}
		return result
	}
	for i := 0; i < 10; i++ {
		if rand.Intn(2) == 1 {
			result += fmt.Sprintf("%v-", RandomInt(1, 50))
			continue
		}
		if rand.Intn(2) == 1 {
			result += RandomString(RandomInt(1, 10))
		}
		if rand.Intn(2) == 1 {
			result += "-"
		}
	}
	return result
}

func RandomString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	str := make([]rune, n)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(str)
}

func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getDigit(v string) (uint, error) {
	digit, err := strconv.Atoi(v)
	if err == nil {
		return uint(digit), nil
	}
	return 0, err
}
