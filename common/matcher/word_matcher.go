package matcher

import (
	"log"
	"regexp"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/reiver/go-porterstemmer"
)

func normalize(text string) []string {

	log.Println("Original:", text)

	text = strings.ToLower(text)

	text = stopwords.CleanString(text, "en", true)
	log.Println("After stopword removal:", text)

	re := regexp.MustCompile(`[^a-z\s]+`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)

	var stemmed []string
	for _, word := range words {
		stem := porterstemmer.StemString(word)
		stemmed = append(stemmed, stem)
	}

	log.Println("After stemming:", stemmed)
	log.Println("---------------")

	return stemmed
}

func MatchPercent(input, target string) float64 {

	log.Println("=== INPUT NORMALIZATION ===")
	inputWords := normalize(input)

	log.Println("=== TARGET NORMALIZATION ===")
	targetWords := normalize(target)

	if len(targetWords) == 0 {
		return 0
	}

	inputSet := make(map[string]struct{})
	for _, word := range inputWords {
		inputSet[word] = struct{}{}
	}

	matches := 0
	var matchedWords []string

	for _, word := range targetWords {
		if _, exists := inputSet[word]; exists {
			matches++
			matchedWords = append(matchedWords, word)
		}
	}

	log.Println("Matched words:", matchedWords)
	log.Printf("Matches: %d / %d\n", matches, len(targetWords))

	return (float64(matches) / float64(len(targetWords))) * 100
}