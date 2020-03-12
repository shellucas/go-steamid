package main

import (
	"fmt"
	"steamid"
)

func main() {
	sid := steamid.CreateSteamID("STEAM_0:1:188656116")
	fmt.Println(sid.Steam2())

	sid = steamid.CreateSteamID("[U:1:377312233]")
	fmt.Println(sid.Steam2())

	sid = steamid.CreateSteamID("76561198337577961")
	fmt.Println(sid.Steam2())
	fmt.Println(sid.ToString())

}
