package main

import (
	"fmt"

	"github.com/sagottlieb/hackerrank/balancedBraces/mystack"
)

func main() {

	bracketDict := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	// checker := basic.NewBalanceChecker(bracketDict)

	checker := mystack.NewBalanceChecker(bracketDict)

	inputs := parseFromStdin()

	for _, s := range inputs {
		if checker.IsBalanced(s) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

type balanceChecker interface {
	isBalanced(string) bool
}
