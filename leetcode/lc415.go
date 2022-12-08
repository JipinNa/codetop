package leetcode

func AddStrings(num1 string, num2 string) string {
	return addStrings(num1, num2)
}

func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	i := len1 - 1
	j := len2 - 1
	var step bool
	var reverseAns string
	for i >= 0 || j >= 0 || step {
		var num uint8
		if i != -1 {
			num += num1[i] - '0'
			i--
		}
		if j != -1 {
			num += num2[j] - '0'
			j--
		}
		if step {
			step = false
			num++
		}
		if num >= 10 {
			num -= 10
			step = true
		}
		reverseAns += string(num + '0')
	}
	var ans string
	for in := len(reverseAns) - 1; in >= 0; in-- {
		ans += string(reverseAns[in])
	}
	return ans
}
