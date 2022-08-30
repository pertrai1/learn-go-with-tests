package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func Sum(nums []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(nums, add, 0)
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sums, []int{})
}

func Reduce[A any](collection []A, accumulator func(A, A) A, initialValue A) A {
	var result = initialValue

	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
