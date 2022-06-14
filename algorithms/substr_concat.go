package algorithms

import (
	"log"
	"strings"
)

var text string
var words []string
var results []int

func substr_concat(t string, ws []string) []int {
	log.Printf("Params::: input text: %s, words: %v", t, ws)
	words = ws
	new_size := len(words)
	text = t
	words = ws
	return substr_concat_helper(new_size)
}

func substr_concat_helper(new_size int) []int {
	log.Printf("Current size: %d", new_size)
	if new_size == 1 {
		return []int{}
	}
	for idx := 0; idx < new_size; idx++ {
		substr_concat_helper(new_size - 1)
		if new_size == 2 {
			concated_words := ""
			for _, word := range words {
				concated_words += word
			}
			log.Printf("Concated words: %s", concated_words)
			results = append(results, matched_position(text, concated_words)...)
		}
		rotate_words(words, new_size)
	}
	return results
}

func rotate_words(w []string, size int) []string {
	start := len(w) - size
	log.Printf("Words: %v, size: %d, start: %d", w, size, start)
	if size < 2 {
		return w
	}
	idx := start
	w_0 := w[idx]
	log.Printf("start word: %s", w_0)
	for idx = start; idx < len(w)-1; idx++ {
		log.Printf("word[idx]: %s, idx: %d", w[idx], idx)
		w[idx] = w[idx+1]
		log.Printf("start word: %s", w[idx])
	}
	w[idx] = w_0
	return w
}

func matched_position(text string, match_str string) []int {
	result := []int{}
	prev_pos := 0
	l := len(match_str)
	for len(text) > l {
		log.Printf("Text: %s, prev_pos: %d", text, prev_pos)
		pos := strings.Index(text, match_str)
		if pos > -1 {
			prev_pos += pos
			result = append(result, prev_pos)
			text = text[pos+l:]
		} else {
			break
		}
	}

	for i, _ := range result {
		result[i] = result[i] + l*i
	}

	return result
}
