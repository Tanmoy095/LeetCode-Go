package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapST := make(map[rune]rune)
	mapTS := make(map[rune]rune)

	for i := range s {
		cs := rune(s[i])
		ct := rune(t[i])

		// Check forward mapping
		if val, ok := mapST[cs]; ok {
			if val != ct {
				return false
			}
		} else {
			mapST[cs] = ct
		}

		// Check reverse mapping
		if val, ok := mapTS[ct]; ok {
			if val != cs {
				return false
			}
		} else {
			mapTS[ct] = cs
		}
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
	fmt.Println(isIsomorphic("ab", "aa"))       // false
}
