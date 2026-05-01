// Package twosum solves the classic Two Sum problem (LeetCode #1).
package twosum

// TwoSum returns the indices of the two numbers in nums that add up to target.
// It returns nil if no such pair exists.
//
// Time:  O(n) — single pass.
// Space: O(n) — the lookup map.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
