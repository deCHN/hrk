package main_test

import "testing"

//https://www.hackerrank.com/contests/moodys-analytics-2018-university-codesprint/challenges/short-trade-transaction/problem

/*
* You need to borrow exactly _a_ units of shares from the market. To do so, you
* visit _m_ market participants one by one and ask them to lend you the remaining
* amount of shares you need.
* Each visited participant can lend you any unit of shares ranging from 0 to
* max(0, min(x-1)), both inclusive, where _x_ is the number of shares lent by
* the previously visited participant or _unlimit_ if no participant has been yet
* visited.

* In how many different ways can the _m_ market participants lend you exactly _a_
* units of shares? Your task is to answer _q_ such queries.
 */
func TestShortTradeTransaction(t *testing.T) {
	tests := []struct {
		a, m int
		want int
	}{
		{2, 2, 1},
		{3, 2, 2},
	}

	for _, v := range tests {
		if get := shortTradeTransaction(v.a, v.m); get != v.want {
			t.Errorf("Given: %v, want: %v, but: %v.\n", v, v.want, get)
		}
	}
}

func shortTradeTransaction(a, m int) int {
	return 0
}
