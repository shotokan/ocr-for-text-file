package ocr


var MAX_LINES_PER_ACOUNT = 3

// Split and transform lines of text in accounts form by 3 lines
// Each account is formed by 9 numbers and each number is a 3x3 characters
func ParseAccounts(rawAccounts []string) [][]string{
	var accountByLines []string = make([]string, 0, MAX_LINES_PER_ACOUNT)
	numberAccounts := (len(rawAccounts)/4) + 1
	
	accounts := make([][]string, 0, numberAccounts)

	for _, textLine := range rawAccounts {
	
		if len(accountByLines) == MAX_LINES_PER_ACOUNT {
			accounts = append(accounts, accountByLines)
			accountByLines = []string{}
			continue
		}
		accountByLines = append(accountByLines, textLine)
		
	}
	return accounts
}

