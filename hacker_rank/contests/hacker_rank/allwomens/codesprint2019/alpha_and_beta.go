package codesprint2019

func alphaBeta(pile []int32) int32 {
	var alpha, beta int32
	for _, cur := range pile {
		if cur > alpha {
			alpha, beta = cur, alpha
		} else if cur < alpha && cur > beta {
			beta = cur
		}
	}
	return beta
}
