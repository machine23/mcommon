package mcommon

import (
	"regexp"
	"sort"
)

var re = regexp.MustCompile(`\w+`)

// MostCommon10 gets a text string and returns a list of the 10 most common words.
func MostCommon10(text string) []string {
	wordMap := make(map[string]int)
	words := re.FindAllString(text, -1)
	for _, w := range words {
		wordMap[w]++
	}
	result := []string{}
	sorted := sortWordByQuantity(wordMap)
	for _, v := range sorted {
		result = append(result, v.word)
		if len(result) == 10 {
			return result
		}
	}
	return result
}

type wordStat struct {
	word     string
	quantity int
}

type wordList []wordStat

func (p wordList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p wordList) Len() int           { return len(p) }
func (p wordList) Less(i, j int) bool { return p[i].quantity > p[j].quantity }

func sortWordByQuantity(m map[string]int) wordList {
	p := make(wordList, len(m))
	i := 0
	for word, quantity := range m {
		p[i] = wordStat{word, quantity}
		i++
	}
	sort.Sort(p)
	return p
}
