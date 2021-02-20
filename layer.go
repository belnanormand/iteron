package iteron

//Layer Layers are used to draw sprites on different Z axis
type Layer struct {
	Name    string
	Order   int
	sprites []*Sprite
	scene   *Scene
}

//NewLayer Creates a new layer, adds it the the scene object and returns
//the created layer for setup purpouses
// name: reference name of the scene
// order: draw order of the layer
// scene: scene the layer will be added to
func NewLayer(name string, order int, scene *Scene) *Layer {

	l := Layer{
		Name:    name,
		Order:   order,
		scene:   scene,
		sprites: make([]*Sprite, 0, 10), //TODO: make capasity customizable
	}

	l.scene.layers[name] = &l
	return &l
}

//AddSprite Adds a sprite to the layer. Composite sprites
//can be added by adding the main sprite only
func (l *Layer) AddSprite(sprite *Sprite) {
	l.sprites = append(l.sprites, sprite)
	//TODO Set the sprites current active layer

}

//Update updates the layer trees
// dt: delta time
func (l *Layer) Update(dt float64) {
	//TODO: Add sprite update functins
}

//Draw draws the layer tree
func (l *Layer) Draw() {
	//TODO: Add sprite draw function
}
