package iteron

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Game The main game entry point
type Game struct {
	window         *pixelgl.Window
	scenes         map[string]*Scene
	currentScene   *Scene
	imageResources map[string]*pixel.Picture
}

//NewGame Create and initialize a game object
// win: pixel windo that will be used for drawing
func NewGame(win *pixelgl.Window) Game {
	game := Game{
		window:         win,
		scenes:         make(map[string]*Scene),
		imageResources: make(map[string]*pixel.Picture),
	}
	return game
}

//Update updates the game state and draws the scenes to the window
func (g *Game) Update() {
	last := time.Now()

	for !g.window.Closed() {

		dt := time.Since(last).Seconds()
		last = time.Now()

		if g.currentScene != nil {
			g.currentScene.Update(dt)
			g.currentScene.Draw()
		}

		g.window.Update()

	}
}

//TransitionToScene transitions to the given scene
// name : name reference of the scene to transition to
func (g *Game) TransitionToScene(name string) {
	g.currentScene = g.scenes[name]
	g.currentScene.Prepare()
}

//LoadImageResource Loads an image into the game image resources.
// path: path to the image resource
// name: the referenc name for the image
func (g *Game) LoadImageResource(path string, name string) {
	resource, err := loadPicture(path)
	if err != nil {
		log.Fatalf("Unable to load resource %v\n%v\n", path, err)
	}
	g.imageResources[name] = &resource
}

//GetImageResource Retrieves a loaded image resource
func (g *Game) GetImageResource(name string) *pixel.Picture {
	return g.imageResources[name]
}
