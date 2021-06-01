package taglessSwitch

var fallbackCaseOne = false
var fallbackCaseTwo = false

func switchMe(a,b bool) {
	switch {
	case a:
		fallbackCaseOne = true
	case b:
		fallbackCaseTwo = true
	}
}
