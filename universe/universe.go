package universe

type Universe int

const (
	// INVALID steamID is Invalid
	INVALID Universe = 0
	// PUBLIC steamID is Public
	PUBLIC Universe = 1
	// BETA steamID is Beta
	BETA Universe = 2
	// INTERNAL steamID is Internal
	INTERNAL Universe = 3
	// DEV steamID is Dev
	DEV Universe = 4
)

// func GetUniverse(nr int) Universe {
// 	if nr == 1 {
// 		return PUBLIC
// 	} else if nr == 2 {
// 		return BETA
// 	} else if nr == 3 {
// 		return INTERNAL
// 	} else if nr == 4 {
// 		return INTERNAL
// 	} else if nr == 5 {
// 		return DEV
// 	} else {
// 		return INVALID
// 	}
// }
