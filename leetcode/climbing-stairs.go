package main

// https://leetcode.com/problems/climbing-stairs/

var myMap = make(map[int]int)

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	} else {
		if _, isFound := myMap[n]; isFound {
			return myMap[n]
		}
		res := climbStairs(n-1) + climbStairs(n-2)
		myMap[n] = res
		return res
	}
}
