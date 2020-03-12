package chatinstanceflag

type ChatInstanceFlags int

const (
	Clan     ChatInstanceFlags = (0x000FFFFF + 1) >> 1
	Lobby    ChatInstanceFlags = (0x000FFFFF + 1) >> 2
	MMSLobby ChatInstanceFlags = (0x000FFFFF + 1) >> 3
)
