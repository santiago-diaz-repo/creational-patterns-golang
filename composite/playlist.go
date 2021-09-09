package composite

import "fmt"

type Playlister interface {
	Component
	AddComponent(Component)
	GetComponents() []Component
}
type Playlist struct {
	Components []Component
	Name       string
}

func NewPlaylist(name string) Playlister {
	return &Playlist{
		Components: make([]Component, 0, 10),
		Name:       name,
	}
}

func (p *Playlist) Play() {
	fmt.Println("Play from playlist " + p.Name)
}

func (p *Playlist) Stop() {
	fmt.Println("Stop from playlist " + p.Name)
}

func (p *Playlist) getName() string {
	return p.Name
}

func (p *Playlist) AddComponent(component Component) {
	p.Components = append(p.Components, component)
	fmt.Println("")
}

func (p *Playlist) RemoveComponent(component Component) {
	for i, v := range p.Components {
		if v.getName() == component.getName() {
			p.Components = append(p.Components[:i], p.Components[i+1:]...)
		}
	}
}

func (p *Playlist) GetComponents() []Component {
	return p.Components
}
