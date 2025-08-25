package useless

import (
	"math/rand"
	"slices"
)

type Randomizer struct{}

func (r Randomizer) Process(input string) (output string, err error) {
	result := randomizeString(input)
	return result, nil
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

func createCharacterSet(s string) []rune {
	runes := []rune(s)
	slices.Sort(runes)
	unique := slices.Compact(runes)
	return unique
}
