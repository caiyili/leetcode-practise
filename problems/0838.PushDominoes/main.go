package leetcode

func pushDominoes(dominoes string) string {
	if dominoes == "" {
		return ""
	}
	var length = len(dominoes)
	var result = make([]byte, length)
	var start = -1
	for i := 0; i < len(dominoes); i++ {
		result[i] = dominoes[i]
		if start < 0 {
			start = i
		}
		cur := dominoes[i]
		switch cur {
		case 'L':
			handle(result[start : i+1])
			start = -1
		case 'R':
			handle(result[start:i])
			start = i
		case '.':
			continue
		default:
			return ""
		}
	}
	if start >= 0 {
		handle(result[start:length])
	}
	return string(result)
}

func pushAllLeft(dominoes []byte) {
	for i := 0; i < len(dominoes); i++ {
		dominoes[i] = 'L'
	}
}

func pushBothEnd(dominoes []byte) {
	var length = len(dominoes)
	if length < 1 {
		return
	}
	for i, j := 0, length-1; i < j; {
		dominoes[i] = 'R'
		dominoes[j] = 'L'
		i++
		j--
	}
}

func pushAllRight(dominoes []byte) {
	for i := 0; i < len(dominoes); i++ {
		dominoes[i] = 'R'
	}
}

func handle(dominoes []byte) {
	var length = len(dominoes)
	if length == 0 {
		return
	}
	var left, right = dominoes[0], dominoes[length-1]
	if left == 'R' && right == 'L' {
		//往中间倒
		pushBothEnd(dominoes)
	} else if left == 'R' {
		//都往右边倒
		pushAllRight(dominoes)
	} else if right == 'L' {
		//都往左边倒
		pushAllLeft(dominoes)
	} else {
		//不倒
	}
}
