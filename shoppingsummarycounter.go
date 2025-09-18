package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	groceries := strings.Split(str, " ")
	receiptCounter := make(map[string]int)
	for _, gr := range groceries {
		receiptCounter[gr]++
	}
	return receiptCounter
}
