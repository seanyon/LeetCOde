/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
import "math"
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

/*
simple solution
func reverse(x int) int {
	minInt32, maxInt32 := -2147483648, 2147483647
	result := 0
	for ; x != 0; x = x/10 {//x=0でなければ, xはint型なので自動で切り捨て
			tempResult := (result * 10) + x % 10//負の値のあまりは負で出力してくれる
			if tempResult > maxInt32 || tempResult < minInt32  {//||または
					return 0
			}
			result = tempResult
	}
	return result
}
*/