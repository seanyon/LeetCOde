/*
Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.



Example 1:

Input: x = 121
Output: true
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
*/

func reverse(x int) int {
	minRange, maxRange := int(-math.Pow(2, 31)), int(math.Pow(2, 31)-1)
	result := 0
	for ; x != 0; x = x / 10 {
		tempResult := result*10 + x%10
		if tempResult < minRange || tempResult > maxRange {
			return 0
		}
		result = tempResult
	}
	return result
}

func isPalindrome(x int) bool {

	if reverse(x) == x && x >= 0 {
		return true
	} else {
		return false
	}

}