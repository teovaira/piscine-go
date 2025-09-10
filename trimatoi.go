package piscine

func TrimAtoi(s string) int {
	sign := 1        
	number := 0         
	foundDigit := false 

	for _, r := range s {         
		if !foundDigit && r == '-' {
			
			sign = -1
			continue
		}

		if r >= '0' && r <= '9' {           
			foundDigit = true
			number = number*10 + int(r-'0')        
		}
		
	}

	if !foundDigit {
		return 0 
	}
	return sign * number
}