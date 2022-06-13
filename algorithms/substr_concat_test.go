package algorithms

import (
	"log"
	"testing"
)

func TestSubstrConcat(t *testing.T) {
	log.Printf("TestSubstrConcat")
	Input := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	// Input = "wordgoodgoodgoodbestword"
	// words = []string{"word", "good", "best", "word"}

	// Input = "barfoofoobarthefoobarman"
	// words = []string{"bar", "foo", "the"}

	result := substr_concat(Input, words)
	log.Printf("Results: %v", result)
}

func TestSubstrConcatRotation(t *testing.T) {
	words := []string{"three", "one", "two"}
	log.Printf("TestSubstrConcatRotation for words: %v", words)
	rotated_words_result := rotate_words(words, 3)
	log.Printf("rotated_words_result: %v", rotated_words_result)
}

func TestMatchedPosition(t *testing.T) {
	match_str := "ber"
	text := "barfoofoobarthefoobarman"
	log.Printf("text: %s, match string: %s", text, match_str)
	result := matched_position(text, match_str)
	log.Printf("Matched Positions: %v", result)
}

/**
You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.

You can return the answer in any order.



Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]


Constraints:

1 <= s.length <= 104
s consists of lower-case English letters.
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] consists of lower-case English letters.
**/
