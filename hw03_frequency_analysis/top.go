package hw03frequencyanalysis

import (
	"math"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

type contWC struct {
	wordCounts []wordCount
}

func (c contWC) Len() int {
	return len(c.wordCounts)
}

func (c contWC) Less(i, j int) bool {
	if c.wordCounts[i].count == c.wordCounts[j].count {
		return strings.Compare(c.wordCounts[i].word, c.wordCounts[j].word) < 0
	}
	return c.wordCounts[i].count > c.wordCounts[j].count
}

func (c contWC) Swap(i, j int) {
	c.wordCounts[i], c.wordCounts[j] = c.wordCounts[j], c.wordCounts[i]
}

func (c contWC) getTopWords() []string {
	min := int(math.Min(float64(len(c.wordCounts)), 10))
	words := make([]string, min)
	for i := 0; i < min; i++ {
		words[i] = c.wordCounts[i].word
	}
	return words
}

func Top10(input string) []string {
	words := strings.Fields(input)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	wordCounts := make([]wordCount, 0, len(m))
	for key, value := range m {
		wordCounts = append(wordCounts, wordCount{word: key, count: value})
	}
	container := contWC{wordCounts: wordCounts}
	sort.Sort(container)
	// Place your code here.
	return container.getTopWords()
}
