package steamid_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shellucas/go-steamid/steamid"
)

func TestSteamIDs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SteamID Suite")
}

var _ = Describe("SteamID", func() {

	Context("STEAM_0:0:458887262", func() {
		sid, err := steamid.CreateSteamID("STEAM_0:0:458887262")

		It("Should not throw an error", func() {
			Expect(err).To(BeNil())
		})

		It("SteamID2 should be the same", func() {
			Expect(sid.GetSteam2RenderedID()).To(Equal("STEAM_0:0:458887262"))
		})

		It("SteamID3 should be converted", func() {
			Expect(sid.GetSteam3RenderedID()).To(Equal("[U:1:917774524]"))
		})

		It("SteamID64 should be converted", func() {
			Expect(sid.GetSteamID64()).To(Equal("76561198878040252"))
		})
	})
})
