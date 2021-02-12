package strings

import "strings"

func morganAndString(a, b string) string {
	var i, j int
	output := strings.Builder{}
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			output.WriteByte(a[i])
			i++
			continue
		}

		if a[i] > b[j] {
			output.WriteByte(b[j])
			j++
			continue
		}

		// 동률일 경우 다음 자리 비교
		if compare(a, b, i+1, j+1) {
			i = extend(a, &output, i)
			continue
		}

		j = extend(b, &output, j)
	}

	if i < len(a) {
		output.WriteString(a[i:])
	}

	if j < len(b) {
		output.WriteString(b[j:])
	}

	return output.String()
}

func extend(in string, out *strings.Builder, i int) int {
	out.WriteByte(in[i])
	i++

	// 반복 문자열은 일괄 붙이기
	for i < len(in) && in[i] == in[i-1] {
		out.WriteByte(in[i])
		i++
	}

	return i
}

func compare(a, b string, i, j int) bool {
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			return true
		}

		if a[i] > b[j] {
			return false
		}

		i++
		j++
	}

	return i != len(a)
}
