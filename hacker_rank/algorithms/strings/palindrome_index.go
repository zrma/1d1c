package strings

import (
	"1d1go/utils/str"
)

func equalPrefix(s1, s2 string) bool {
	eq := s1[0] == s2[0]
	if len(s1) < 2 || len(s2) < 2 {
		return eq
	}

	return eq && (s1[1] == s2[1])
}

type checker func(int) (target int, isDiff, available bool)

func diffChecker(s1, s2 string, length int) checker {
	var forward, backward bool
	var forwardMargin, backwardMargin, target int

	isDiff := func() bool {
		return forward || backward
	}

	// 현재 위치에서 1글자만 더 바꾸면 회문이 될 수 있는지 검증하는 함수
	// closure 로 isDiff 값을 저장하고 있다(이미 한 번 기회를 소진했는지 여부)
	return func(curPos int) (int, bool, bool) {
		// 회문이다.
		if s1[curPos+forwardMargin] == s2[curPos+backwardMargin] {
			return target, isDiff(), true
		}

		// 딱 한 글자만 바꿀 수 있는데, 그 기회를 이미 소모했을 경우.
		if isDiff() {
			return -1, isDiff(), false
		}

		// 아래의 두 분기는 모두 동시적으로 처리해야 한다.(순차적으로 처리하면 안 됨)
		// 그러므로 한꺼번에 다음 글자까지(equalPrefix 함수를 사용하여) 체크해야 한다.

		// 만약 역주행 방향( <- )으로 한 글자를 바꿔서 같아질 수 있다면(그리고 그 다음 글자도 같다면)
		// 역주행 쪽에 표시 하고 마진을 1 추가.
		if s1[curPos] == s2[curPos+1] && equalPrefix(s1[curPos:], s2[curPos+1:]) {
			backwardMargin++
			backward = true
			target = length - curPos - 1

			return target, isDiff(), true
		}

		// 만약 정주행 방향( -> )으로 한 글자를 바꿔서 같아질 수 있다면(그리고 그 다음 글자도 같다면)
		// 정주행 쪽에 표시하고 마진을 1 추가.
		if s1[curPos+1] == s2[curPos] && equalPrefix(s1[curPos+1:], s2[curPos:]) {
			forwardMargin++
			forward = true
			target = curPos

			return target, isDiff(), true
		}

		return -1, isDiff(), false
	}
}

func diffPos(s1, s2 string, length int) (target int, diff bool) {

	diffChecker := diffChecker(s1, s2, length)
	var valid bool
	for i := 0; i < length/2; i++ {
		if target, diff, valid = diffChecker(i); !valid {
			// 한 글자를 바꾸더라도 회문이 될 수 없다.
			return -1, diff
		}
	}

	return target, diff
}

func palindromeIndex(s string) int32 {
	length := len(s)
	if length < 2 {
		return -1
	}

	re := str.Reverse(s)
	pos, diff := diffPos(s, re, length)
	if !diff {
		// 만약 이미 회문이라면 더 이상 제거할 문자가 없다
		return -1
	}

	return int32(pos)
}
