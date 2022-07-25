package algorithms

import (
	"log"
	"strings"
)

func text_justification(words []string, maxWidth int) []string {

	results := make([]string, 1)

	results[0] = ""
	return text_justification_helper(words, maxWidth, results)
}

func text_justification_helper(words []string, maxWidth int, results []string) []string {
	if len(words) < 1 {
		// results[len(results)-1] = format_text(results[len(results)-1], maxWidth)
		if maxWidth-len(results[len(results)-1])-1 > 0 {
			results[len(results)-1] = results[len(results)-1] + strings.Repeat(" ", maxWidth-len(results[len(results)-1]))
		}
		return results
	}
	log.Printf("Words: %v, maxWidth: %d, results: %v", words, maxWidth, results)

	last_item := results[len(results)-1]
	log.Printf("The current last item: %s", last_item)
	if len(words[0]+" "+last_item)-1 > maxWidth {
		// results[len(results)-1] = format_text(last_item[0:len(last_item)-1], maxWidth)
		results[len(results)-1] = format_text(last_item, maxWidth)
		results = append(results, "")
		log.Printf("Length of string array: %d", len(results))
		log.Printf("Results: %v", results)
		log.Printf("The last item: %s is exceedind the maxWidth: %d when adding: %s, Overflow is handled by adding a new item", results[len(results)-1], maxWidth, words[0])
		return text_justification_helper(words, maxWidth, results)
	} else {
		last_item += words[0] + " "
		results[len(results)-1] = last_item
		log.Printf("Added the word: %s to the result: %s", words[0], results[len(results)-1])
		return text_justification_helper(words[1:], maxWidth, results)
	}
	return results
}

func format_text(input string, maxWidth int) string {
	l := len(input)
	d := maxWidth - l
	if d < 1 {
		log.Printf("The max width is bigger than the given input")
		return input
	}
	words := strings.Fields(input)
	nw := len(words)
	if nw == 1 {
		return input + strings.Repeat(" ", d)
	}
	n_s := int(d / (nw - 1))
	log.Printf("input: %s, input length: %d, diff: %d, maximum width: %d, number of words: %d, number of spaces: %d", input, l, d, maxWidth, nw, n_s)

	spces := strings.Repeat(" ", n_s)
	rem_spces := maxWidth
	for idx, _ := range words[0 : len(words)-1] {
		words[idx] = words[idx] + " " + spces
		rem_spces -= len(words[idx])
	}
	rem_spces -= len(words[len(words)-1])
	for idx := 0; idx < rem_spces; idx++ {
		words[idx] += " "
	}
	log.Printf("Remaining spaces: %d", rem_spces)
	return strings.Join(words, "")
}
