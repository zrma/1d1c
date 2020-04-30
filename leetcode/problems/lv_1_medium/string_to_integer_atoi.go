package lv_1_medium

import (
	"math"
	"strconv"
	"strings"
)

func myAtoI(str string) int {
	str = strings.TrimLeft(str, " ")
	if str == "" {
		return 0
	}

	var num int64
	var negative bool
	if str[0] == '-' {
		negative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	str = strings.TrimLeft(str, "0")
	if str == "" {
		return 0
	}

Loop:
	for i, c := range str {
		if i > 10 {
			break
		}
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			cur, _ := strconv.Atoi(string(c))
			num = num*10 + int64(cur)
		default:
			break Loop
		}
	}

	if negative {
		num = -num
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		num = math.MinInt32
	}
	return int(num)
}
