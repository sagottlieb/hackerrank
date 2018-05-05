package main

import (
	"fmt"
	"github.com/sagottlieb/hackerrank/stacks/balancedBraces/basic"
)

func main() {

	bracketDict := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	basicChecker := basic.NewBalanceChecker(bracketDict)

	inputs := parseFromStdin()

	for _, s := range inputs{
		if basicChecker.IsBalanced(s){
			fmt.Println("YES")
		} else{
		fmt.Println("NO")
		}
	}

}

type balanceChecker interface{
	isBalanced(string) bool
}
