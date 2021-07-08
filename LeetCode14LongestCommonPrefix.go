/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		s := strs[0][i]
		for _, str := range strs[1:] {
			if i == len(str) || s != str[i] {
				return strs[0][:i]
			}
		}

	}
	return strs[0]
}