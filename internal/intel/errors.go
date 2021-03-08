package intel

import "fmt"

var (
	ErrPatternInvalid  = "Pattern is invalid Go Regexp pattern"
	ErrPatternInvalidf = fmt.Sprint(ErrPatternInvalid, ", %s")
)
