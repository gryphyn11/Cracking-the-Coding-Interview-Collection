package chapter1

import (
	"sort"
)

/*
Scans the runes the given string consists of and checks for uniqueness.
Returns true if all runes in the string are unique, false otherwise.
*/
func IsUniqueRunes(s string) bool {
	runeMap := make(map[rune]bool)
	for _, rune := range s {
		if runeMap[rune] {
			return false
		}
		runeMap[rune] = true
	}
	return true
}

/*
Scans the runes the given string consists of and checks for uniqueness.
This implementation satisfies the requested constraint that data structures not be used by
reslicing the string as a rune slice and sorting it, guaranteeing
equivalent runes will be adjacent. If any pair of adjacent runes match, returns false,
else returns true.
*/
func IsUniqueRunes2(s string) bool {
	runeSlice := []rune(s)
	sort.Sort(runeSorter(runeSlice))
	for i, j := 1, 0; i < len(runeSlice); i, j = i+1, j+1 {
		if runeSlice[i] == runeSlice[j] {
			return false
		}
	}
	return true
}

/*
Type runeSorter is a rune slice that implements sort.Interface.
Doing so allows the slice to be sorted with sort.Sort().
*/
type runeSorter []rune

func (r runeSorter) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSorter) Len() int {
	return len(r)
}

func (r runeSorter) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
