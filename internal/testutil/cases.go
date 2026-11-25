// Package testutil содержит только входы для проверки домашних задач.
package testutil

import "math"

type Case struct {
	Name   string
	Values []int
}

func Cases() []Case {
	return []Case{
		{"nil", nil}, {"empty", []int{}}, {"zero", []int{0}}, {"positive", []int{7}},
		{"negative", []int{-3}}, {"two_ascending", []int{1, 2}}, {"two_reverse", []int{2, 1}},
		{"equal_pair", []int{4, 4}}, {"duplicates", []int{3, 1, 3, 2, 1}},
		{"negatives", []int{-9, -1, -5}}, {"sorted", []int{-2, 0, 3, 5, 8}},
		{"reverse", []int{8, 5, 3, 0, -2}}, {"mixed", []int{-2, 5, 0, 5, -2}},
		{"extremes", []int{math.MinInt, 0, math.MaxInt}}, {"many", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
	}
}
