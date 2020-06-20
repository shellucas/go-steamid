package steamid

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"

	chat "github.com/BlackwonderTF/go-steamid/chatinstanceflag"
	"github.com/BlackwonderTF/go-steamid/instance"
	steamIDType "github.com/BlackwonderTF/go-steamid/type"
	"github.com/BlackwonderTF/go-steamid/universe"
)

type steamID struct {
	universe  universe.Universe
	idType    steamIDType.Type
	instance  instance.Instance
	accountid int
}

// Source: https://developer.valvesoftware.com/wiki/SteamID

// CreateSteamID creates a steamID object
func CreateSteamID(input string) (steamID, error) {
	s := steamID{
		universe:  universe.INVALID,
		idType:    steamIDType.INVALID,
		instance:  instance.ALL,
		accountid: 0,
	}

	if len(input) == 0 {
		return s, errors.New("There was no input given")
	}

	// r := regexp.MustCompile(`STEAM_(?P<Universe>[0-5]):(?P<IDType>[0-1]):(?P<ID>[0-9]+)`)
	isSteamID2 := regexp.MustCompile(`^STEAM_([0-5]):([0-1]):([0-9]+)$`)
	isSteamID3 := regexp.MustCompile(`^\[([a-zA-Z]):([0-5]):([0-9]+)(:[0-9]+)?\]$`)

	if isSteamID2.MatchString(input) {
		matches := isSteamID2.FindStringSubmatch(input)
		iUniverse, _ := strconv.Atoi(matches[1])
		accountid1, _ := strconv.Atoi(matches[3])
		accountid2, _ := strconv.Atoi(matches[2])

		// universe.GetUniverse(iUniverse)
		if iUniverse > 0 {
			s.universe = universe.Universe(iUniverse)
		} else {
			s.universe = universe.PUBLIC
		}

		s.idType = steamIDType.INDIVIDUAL
		s.instance = instance.DESKTOP

		s.accountid = (accountid1 * 2) + accountid2
	} else if isSteamID3.MatchString(input) {
		matches := isSteamID3.FindStringSubmatch(input)
		iUniverse, _ := strconv.Atoi(matches[2])
		iAccountid, _ := strconv.Atoi(matches[3])

		s.universe = universe.Universe(iUniverse)
		s.accountid = iAccountid

		typeChar := matches[1]

		if len(matches[4]) > 0 {
			iInstance, _ := strconv.Atoi(string(matches[4][1]))
			s.instance = instance.Instance(iInstance)
		} else if typeChar == "U" {
			s.instance = instance.DESKTOP
		}

		if typeChar == "C" {
			s.instance = instance.Instance(int(s.instance) | int(chat.Clan))
			s.idType = steamIDType.CHAT
		} else if typeChar == "L" {
			s.instance = instance.Instance(int(s.instance) | int(chat.Lobby))
			s.idType = steamIDType.CHAT
		} else {
			s.idType = steamIDType.GetType(typeChar)
		}
	} else {
		i, _ := strconv.ParseUint(input, 10, 64)
		s.accountid = int(i) & 0xFFFFFFFF >> 0

		i = i >> 32
		s.instance = instance.Instance(int(i) & 0xFFFFF)

		i = i >> 20
		s.idType = steamIDType.Type(int(i) & 0xF)

		i = i >> 4
		s.universe = universe.Universe(int(i))
	}

	return s, nil
}

// FromIndividualAccountID Create an individual SteamID in the public universe given an accountid
func FromIndividualAccountID(accountid int) steamID {
	var sid = steamID{
		universe:  universe.PUBLIC,
		idType:    steamIDType.INDIVIDUAL,
		instance:  instance.DESKTOP,
		accountid: accountid,
	}
	return sid
}

// IsValid Check whether this SteamID is valid (according to Steam's rules)
func (sid steamID) IsValid() bool {
	if sid.idType <= steamIDType.INVALID || sid.idType > steamIDType.ANON_USER {
		return false
	}

	if sid.universe <= universe.INVALID || sid.universe > universe.DEV {
		return false
	}

	if sid.idType == steamIDType.INDIVIDUAL && (sid.accountid == 0 || sid.instance > instance.WEB) {
		return false
	}

	if sid.idType == steamIDType.CLAN && (sid.accountid == 0 || sid.instance != instance.ALL) {
		return false
	}

	if sid.idType == steamIDType.GAMESERVER && sid.accountid == 0 {
		return false
	}

	return true
}

// IsGroupChat Check whether this chat SteamID is tied to a Steam group.
func (sid steamID) IsGroupChat() bool {
	return !!((sid.idType == steamIDType.CHAT) && (int(sid.instance)&int(chat.Clan)) > 0)
}

// IsLobby Check whether this chat SteamID is a Steam lobby.
func (sid steamID) IsLobby() bool {
	return !!((sid.idType == steamIDType.CHAT) && (int(sid.instance)&int(chat.Lobby) > 0 || (int(sid.instance)&int(chat.MMSLobby)) > 0))
}

// Steam2 Render this SteamID into Steam2 textual format
func (sid steamID) Steam2(newerFormat ...bool) string {
	var newFormat bool
	if len(newerFormat) > 0 {
		newFormat = newerFormat[0]
	} else {
		newFormat = false
	}

	if sid.idType != steamIDType.INDIVIDUAL {
		return ""
	} else {
		universe := int(sid.universe)

		if !newFormat && universe == 1 {
			universe = 0
		}

		id := fmt.Sprintf("%0.0f", math.Floor(float64(sid.accountid)/2))

		return fmt.Sprintf("STEAM_%s:%s:%s", strconv.Itoa(universe), strconv.Itoa(sid.accountid&1), id)
	}
}

// GetSteam2RenderedID Render this SteamID into Steam2 textual format
func (sid steamID) GetSteam2RenderedID(newerFormat ...bool) string {
	if len(newerFormat) > 0 {
		return sid.Steam2(newerFormat[0])
	} else {
		return sid.Steam2()
	}
}

// Steam3 Render this SteamID into Steam3 textual format
func (sid steamID) Steam3() string {
	Y := sid.accountid & 1
	Z := int(math.Floor(float64(sid.accountid) / 2))
	W := Z*2 + Y
	return fmt.Sprintf("[%s:1:%d]", sid.idType.GetString(), W)
}

// GetSteam3RenderedID Render this SteamID into Steam3 textual format
func (sid steamID) GetSteam3RenderedID() string {
	return sid.Steam3()
}

// ToString Render this SteamID into 64-bit numeric format
func (sid steamID) ToString() string {
	Z := sid.accountid
	Instance := (int(sid.instance) << 32)
	Type := (int(sid.idType) << 52)
	X := (int(sid.universe) << 56)

	return fmt.Sprintf("%d", Z+Instance+Type+X)
}

// GetSteamID64 Render this SteamID into 64-bit numeric format
func (sid steamID) GetSteamID64() string {
	return sid.ToString()
}
