package composite

import "fmt"

type Songer interface {
	Component
	GetArtist() string
}
type Song struct {
	Name   string
	Artist string
}

func NewSong(name string, artist string) Songer {
	return &Song{
		Name:   name,
		Artist: artist,
	}
}

func (s *Song) Play() {
	fmt.Println("Play from song " + s.Name)
}

func (s *Song) Stop() {
	fmt.Println("Stop from song " + s.Name)
}

func (s *Song) getName() string {
	return s.Name
}

func (s *Song) GetArtist() string {
	return s.Artist
}
