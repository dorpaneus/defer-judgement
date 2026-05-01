# Two Sum

**Source:** [LeetCode #1](https://leetcode.com/problems/two-sum/)
**Difficulty:** easy
**Tags:** `array`, `hashmap`, `lookup`
**Date solved:** 2026-05-01

---

## Problem

Given an array of integers and a target value, return the indices of the two numbers that add up to the target. Each input has exactly one valid answer, and you can't use the same element twice.

**Input:** `nums []int`, `target int`
**Output:** `[]int` of length 2 — the two indices, smaller one first
**Constraints worth noting:** 2 ≤ len(nums) ≤ 10⁴; exactly one valid pair guaranteed by the problem statement.

## Examples

```
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Why:    nums[0] + nums[1] = 2 + 7 = 9
```

## Approach

The brute-force O(n²) double loop works but wastes time. Instead, walk the array once and keep a map from `value → index`. For each `n`, check whether the *complement* `target - n` is already in the map; if it is, we've found the pair.

The trick is that we only need to remember what we've already seen — we never need to look ahead.

## Complexity

- **Time:** O(n) — single pass over the array, O(1) average map ops.
- **Space:** O(n) — at worst the map holds every element before we find a match.

## Lessons / gotchas

- In Go, `make(map[int]int, len(nums))` pre-sizes the map so it doesn't reallocate as it grows. Small habit, costs nothing.
- `reflect.DeepEqual` is the easy way to compare slices in tests — `==` doesn't work on slices in Go.
- Returning `nil` for "not found" is idiomatic; callers can compare against `nil` or check `len(...) == 0`.
- Package name has no hyphen (`twosum`), even though the folder is `two-sum`. Go package names must be a single identifier.

## Related problems

- (none yet — link 3Sum, Two Sum II, etc. when I solve them)
