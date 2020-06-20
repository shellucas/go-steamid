package main

import (
	"log"

	"github.com/BlackwonderTF/go-steamid/steamid"
)

func main() {
	x, _ := steamid.CreateSteamID("STEAM_0:0:458887262")
	y, _ := steamid.CreateSteamID("STEAM_0:1:188656116")

	log.Println(x.Steam3())
	log.Println(y.Steam3())
}
