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

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue

	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	balance := func(current float64, t Transaction) float64 {
		if t.From == name {
			current -= t.Sum
		}
		if t.To == name {
			current += t.Sum
		}
		return current
	}
	return Reduce(transactions, balance, 0.0)
}
