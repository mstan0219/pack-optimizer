package pack

import (
	"maps"
	"sort"
)

var (
	Sizes = []int{5000, 2000, 1000, 500, 250}
)

func SetSizes(newSizes []int) []int {
	Sizes = newSizes
	return Sizes
}

func GetSizes() []int {
	return Sizes
}

// Correct returns a map[size]count covering x using available sizes.
// It greedily fills from largest to smallest, adds one smallest pack if a remainder exists,
// then calls optimizePacks to combine smaller packs into larger ones.
// Precondition: sizes must be sorted descending.
// Example:
//
//	Correct(1,...(sizes:[]int{5000, 2000, 1000, 500, 250}, []int{23, 31, 53})) // -> map[int]int{250:1}
func Correct(x int, sizes []int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	smallest := sizes[len(sizes)-1]
	if x <= smallest {
		return map[int]int{smallest: 1}
	}
	maxDPSize := min(x*2, 1_000_000)
	dp := make([]int, maxDPSize+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	packUsed := make([]int, maxDPSize+1)
	for i := 1; i <= maxDPSize; i++ {
		for _, size := range sizes {
			if i >= size && dp[i-size] != -1 {
				newTotal := dp[i-size] + size
				if dp[i] == -1 || newTotal < dp[i] {
					dp[i] = newTotal
					packUsed[i] = size
				}
			}
		}
	}
	minValid := -1
	for i := x; i <= maxDPSize; i++ {
		if dp[i] != -1 {
			minValid = i
			break
		}
	}
	if minValid == -1 {
		return map[int]int{}
	}
	packs := make(map[int]int)
	remain := minValid
	for remain > 0 {
		size := packUsed[remain]
		packs[size]++
		remain -= size
	}
	optimize(packs, sizes)
	return packs
}

func optimize(packs map[int]int, sizes []int) {
	for i := len(sizes) - 1; i > 0; i-- {
		small := sizes[i]
		large := sizes[i-1]

		if large%small != 0 {
			continue
		}
		requiredSmall := large / small
		if have := packs[small]; have >= requiredSmall {
			convert := have / requiredSmall
			if convert > 0 {
				packs[small] -= convert * requiredSmall
				if packs[small] == 0 {
					delete(packs, small)
				}
				packs[large] += convert
			}
		}
	}
}

// InCorrect returns a list of all incorrect pack combinations for a given ordered amount.
// It generates all possible combinations of packs, calculates the correct combination,
// and then filters out the correct one from the list of all combinations.
// Example:
//
//	InCorrect(1,...(sizes:[]int{5000, 2000, 1000, 500, 250}, []int{23, 31, 53})) // -> []map[int]int{{500:1}, {250:2}, {1000:1}, ...}
//
// Note: This function assumes that the pack sizes are sorted in descending order.
// It generates combinations based on the available pack sizes and the ordered amount.
// It returns a slice of maps, where each map represents a combination of pack sizes and their counts.
// The function uses a greedy approach to find the correct combination and then filters out that combination from
// the list of all combinations to return only the incorrect ones.package pack
func InCorrect(x int, sizes []int) []map[int]int {
	correct := Correct(x, sizes)
	var incorrect []map[int]int

	for _, size := range sizes {
		count := (x + size - 1) / size
		comb := map[int]int{size: count}

		if !maps.Equal(comb, correct) {
			incorrect = append(incorrect, comb)
		}
		if len(incorrect) >= 3 {
			break
		}
	}

	if len(incorrect) < 3 { // count InCorrect items
		if len(sizes) >= 2 {
			alt := map[int]int{
				sizes[0]: x/sizes[0] + 1,
				sizes[1]: 1,
			}
			if !maps.Equal(alt, correct) {
				incorrect = append(incorrect, alt)
			}
		}
	}

	return incorrect
}
