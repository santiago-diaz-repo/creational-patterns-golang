package composite

import (
	"fmt"
	"testing"
)

func Test_Component(t *testing.T) {
	playlist1 := NewPlaylist("playlist One")
	playlist2 := NewPlaylist("playlist Two")
	song1 := NewSong("Song One", "Me")
	song2 := NewSong("Song Two", "Me")

	playlist2.AddComponent(song2)
	playlist1.AddComponent(playlist2)
	playlist1.AddComponent(song1)

	for _, v := range playlist1.GetComponents() {
		switch v.(type) {
		case Songer:
			tmp := v.(Songer)
			fmt.Println(tmp.GetArtist())
			tmp.Play()
		case Playlister:
			tmp := v.(Playlister)
			tmp.GetComponents()
			tmp.Play()
		}
	}
}
