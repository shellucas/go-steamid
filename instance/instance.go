package instance

type Instance int

const (
	// ALL steamID is Invalid
	ALL Instance = 0
	// DESKTOP steamID is Public
	DESKTOP Instance = 1
	// CONSOLE steamID is Beta
	CONSOLE Instance = 2
	// WEB steamID is Dev
	WEB Instance = 4
)
