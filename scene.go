package iteron

import "image/color"

//Scene a specific scene in the game
type Scene struct {
	Name            string
	Game            *Game
	BackgroundColor color.RGBA
	layers          map[string]*Layer
}

//NewScene Creates a new scene, adds it the the game object and returns
//the created scene for setup purpouses
// name: reference name of the scene
// game: game th scene will be added to
func NewScene(name string, game *Game) *Scene {
	s := Scene{
		Name:   name,
		Game:   game,
		layers: make(map[string]*Layer),
	}

	s.Game.scenes[name] = &s
	return &s
}

//Prepare Prepares a scene for during transition
func (s *Scene) Prepare() {

}

//Update updates the scene trees
// dt: delta time
func (s *Scene) Update(dt float64) {
	for _, layer := range s.layers {
		layer.Update(dt)
	}
}

//Draw draws the scene tree
func (s *Scene) Draw() {
	s.Game.Window.Clear(s.BackgroundColor)
	for _, layer := range s.layers {
		layer.Draw()
	}
	
}
