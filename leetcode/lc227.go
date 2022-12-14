package leetcode

import (
	"math"
	"strconv"
)

func Calculate(s string) int {
	return calculate(s)
}

func calculate(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

func calculate1(s string) int {
	f := ""
	r := ""
	var b byte
	i := 0
	for i = range s {
		if s[i] >= '0' && s[i] <= '9' {
			if b&4 == 4 || b&8 == 8 {
				r += string(s[i])
			} else {
				f += string(s[i])
			}
			continue
		}
		if s[i] == '+' {
			b |= 1
			break
		}
		if s[i] == '-' {
			b |= 2
			break
		}
		if s[i] == '*' || s[i] == '/' {
			if b&4 == 4 {
				fi, _ := strconv.Atoi(f)
				ri, _ := strconv.Atoi(r)
				f = strconv.Itoa(fi * ri)
				b &= math.MaxUint8 - 4
			}
			if b&8 == 8 {
				fi, _ := strconv.Atoi(f)
				ri, _ := strconv.Atoi(r)
				f = strconv.Itoa(fi / ri)
				b &= math.MaxUint8 - 8
			}
			r = ""
			if s[i] == '*' {
				b |= 4
				continue
			}
			if s[i] == '/' {
				b |= 8
				continue
			}
		}
	}
	fi, _ := strconv.Atoi(f)
	if b&4 == 4 {
		ri, _ := strconv.Atoi(r)
		fi = fi * ri
		b &= math.MaxUint8 - 4
	}
	if b&8 == 8 {
		ri, _ := strconv.Atoi(r)
		fi = fi / ri
		b &= math.MaxUint8 - 8
	}
	if b&1 == 1 {
		return fi + calculate(s[i+1:])
	}
	if b&2 == 2 {
		return fi - calculate(s[i+1:])
	}
	return fi
}
