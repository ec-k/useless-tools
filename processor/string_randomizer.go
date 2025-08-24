package useless

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
)

func createCharacterSet(s string) []rune {
	runes := []rune(s)
	slices.Sort(runes)
	unique := slices.Compact(runes)
	return unique
}

func randomizeString(raw_string string) string {
	length := len([]rune(raw_string))
	uniqueStringSet := createCharacterSet(raw_string)

	randomizedString := make([]rune, length)
	for i:=0;i<length;i++{
		runeIndex := rand.Intn(len(uniqueStringSet))
		randomRune := uniqueStringSet[runeIndex]
		randomizedString[i] = randomRune
	}

	return string(randomizedString)
}

func parseArgs() string {
	return os.Args[1]
}

func main(){
	var inputString string = parseArgs()
	fmt.Println(randomizeString(inputString))
}

