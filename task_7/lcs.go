package task7

import "sort"

func Find(s1, s2 string) []string {
	a := []rune(s1)
	b := []rune(s2)

	dp := buildLCSLengthDP(a, b)
	seen := make(map[string]struct{})
	collectLCS(a, b, len(a), len(b), dp, "", seen)

	result := make([]string, 0, len(seen))
	for seq := range seen {
		result = append(result, seq)
	}
	sort.Strings(result)
	return result
}

func buildLCSLengthDP(a, b []rune) [][]int {
	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp
}

func collectLCS(a, b []rune, i, j int, dp [][]int, current string, result map[string]struct{}) {
	if i == 0 || j == 0 {
		result[current] = struct{}{}
		return
	}

	if a[i-1] == b[j-1] {
		collectLCS(a, b, i-1, j-1, dp, string(a[i-1])+current, result)
		return
	}

	if dp[i-1][j] == dp[i][j] {
		collectLCS(a, b, i-1, j, dp, current, result)
	}

	if dp[i][j-1] == dp[i][j] {
		collectLCS(a, b, i, j-1, dp, current, result)
	}
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}
	return right
}
