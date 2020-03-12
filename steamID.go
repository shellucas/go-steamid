package steamid

import (
	"regexp"
	"steamid/instance"
	steamIDType "steamid/type"
	"steamid/universe"
)

type steamID struct {
	universe  universe.Universe
	idType    steamIDType.Type
	instance  instance.Instance
	accountid int
}

func createSteamID(input string) steamID {
	steamID := steamID{
		universe:  universe.INVALID,
		idType:    steamIDType.INVALID,
		instance:  instance.ALL,
		accountid: 0,
	}

	if len(input) == 0 {
		return steamID
	}

	isSteamID2, err := regexp.MatchString("^STEAM_([0-5]):([0-1]):([0-9]+)$", input)

	if err != nil {
		return steamID
	}

	isSteamID3, err := regexp.MatchString("^[([a-zA-Z]):([0-5]):([0-9]+)(:[0-9]+)?]$", input)

	if err != nil {
		return steamID
	}

	if isSteamID2 {
		// TODO
		/*
			this.universe = parseInt(matches[1], 10) || SteamID.Universe.PUBLIC; // If it's 0, turn it into 1 for public
			this.type = SteamID.Type.INDIVIDUAL;
			this.instance = SteamID.Instance.DESKTOP;
			this.accountid = (parseInt(matches[3], 10) * 2) + parseInt(matches[2], 10);
		*/
	} else if isSteamID3 {
		// TODO
		/*
			this.universe = parseInt(matches[2], 10);
			this.accountid = parseInt(matches[3], 10);

			var typeChar = matches[1];

			if (matches[4]) {
				this.instance = parseInt(matches[4].substring(1), 10);
			} else if (typeChar == 'U') {
				this.instance = SteamID.Instance.DESKTOP;
			}

			if (typeChar == 'c') {
				this.instance |= SteamID.ChatInstanceFlags.Clan;
				this.type = SteamID.Type.CHAT;
			} else if (typeChar == 'L') {
				this.instance |= SteamID.ChatInstanceFlags.Lobby;
				this.type = SteamID.Type.CHAT;
			} else {
				this.type = getTypeFromChar(typeChar);
			}
		*/
	} else {
		// TODO
		/*
			var num = new UInt64(input, 10);
			this.accountid = (num.toNumber() & 0xFFFFFFFF) >>> 0;
			this.instance = num.shiftRight(32).toNumber() & 0xFFFFF;
			this.type = num.shiftRight(20).toNumber() & 0xF;
			this.universe = num.shiftRight(4).toNumber();
		*/
	}

	return steamID
}
