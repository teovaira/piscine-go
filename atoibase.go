package piscine

func AtoiBase(s string, base string) int {
	br := []rune(base)
	if len(br) < 2 {
		return 0
	}
	for i := 0; i < len(br); i++ {
		if br[i] == '+' || br[i] == '-' {
			return 0
		}
		for j := i + 1; j < len(br); j++ {
			if br[i] == br[j] {
				return 0
			}
		}
	}

	if len([]rune(s)) == 0 {
		return 0
	}

	b := len(br)
	res := 0

	for _, r := range s {
		d := -1
		for i := 0; i < b; i++ {
			if br[i] == r {
				d = i
				break
			}
		}
		if d == -1 {
			return 0
		}
		res = res*b + d
	}

	return res
}
