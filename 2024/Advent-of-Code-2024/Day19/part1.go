package main

import "fmt"

func solveA(patterns []string, designs []string) {
	possible := 0

	towels := make(map[string]bool)
	for _, pattern := range patterns {
		towels[pattern] = true
	}

	for _, design := range designs {
		n := len(design)
		dp := make([]bool, n+1)
		dp[0] = true

		for i := 1; i <= n; i++ {
			for j := 0; j < i; j++ {
				if dp[j] && towels[design[j:i]] {
					dp[i] = true
					break
				}
			}
		}

		if dp[n] {
			possible++
		}
	}

	fmt.Println(possible)
}
