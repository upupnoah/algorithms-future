package main

// 模拟
// 思路：
// 每次只有 1 种选择：
// 1. 将 pushed 中的下一个数字压入栈中；
// 2. 将栈顶的数字弹出栈，匹配 popped 序列中下一个弹出的数字
// 因此只要栈顶元素与 popped 序列中下一个弹出的数字相等，就应该执行 2，否则执行 1
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	cnt := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[cnt] {
			stack = stack[:len(stack)-1]
			cnt++
		}
	}
	return len(stack) == 0
}
