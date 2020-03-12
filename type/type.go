package steamIDType

type Type int

const (
	// INVALID steamID is Invalid
	INVALID Type = 0
	// INDIVIDUAL steamID is Public
	INDIVIDUAL Type = 1
	// MULTISEAT steamID is Beta
	MULTISEAT Type = 2
	// GAMESERVER steamID is Dev
	GAMESERVER Type = 3
	// ANON_GAMESERVER steamID is Dev
	ANON_GAMESERVER Type = 4
	// PENDING steamID is Dev
	PENDING Type = 5
	// CONTENT_SERVER steamID is Dev
	CONTENT_SERVER Type = 6
	// CLAN steamID is Dev
	CLAN Type = 7
	// CHAT steamID is Dev
	CHAT Type = 8
	// P2P_SUPER_SEEDER steamID is Dev
	P2P_SUPER_SEEDER Type = 9
	// ANON_USER steamID is Dev
	ANON_USER Type = 10
)

func GetType(t string) Type {
	if t == "I" {
		return INVALID
	} else if t == "U" {
		return INDIVIDUAL
	} else if t == "M" {
		return MULTISEAT
	} else if t == "G" {
		return GAMESERVER
	} else if t == "A" {
		return ANON_GAMESERVER
	} else if t == "P" {
		return PENDING
	} else if t == "C" {
		return CONTENT_SERVER
	} else if t == "g" {
		return CLAN
	} else if t == "T" {
		return CHAT
	} else if t == "a" {
		return P2P_SUPER_SEEDER
	} else {
		return INVALID
	}
}
