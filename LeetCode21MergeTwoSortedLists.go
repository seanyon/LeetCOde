/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true
*/

//S型定義
type S struct {
	items []byte
}

//pushメソッド:加える
func (x *S) Push(y byte) {
	x.items = append(x.items, y)
}

//popメソッド:最後を削る
func (x *S) Pop() byte {
	lstindex := len(x.items) - 1
	xlstitem := x.items[lstindex] //一番最後を取得
	x.items = x.items[:lstindex]
	return xlstitem
}

//IsEmptyメソッド:空かどうか判定
func (x *S) IsEmpty() bool {
	return len(x.items) == 0
}
func isValid(s string) bool {
	//stackをS型として定義
	stack := S{}
	//parsをmapで保持
	pars := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//sの文字数分だけ繰り返す
	for i := 0; i < len(s); i++ {
		v, ok := pars[s[i]] //v:s[i]の相方カッコはじめ。ok:カッコ終わりかどうか
		//カッコはじめならstackにpush
		if !ok {
			stack.Push(s[i])
			//カッコ終わりなら
			//空ならflase
		} else if stack.IsEmpty() {
			return false
		} else {
			popped_item := stack.Pop()
			if popped_item != v { //stackから削り、それが対応する始めでないならfalse
				return false
			}

		}
	}
	//最後IsEmptyを返す
	return stack.IsEmpty()

}