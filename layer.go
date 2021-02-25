package iteron

//Layer Layers are used to draw sprites on different Z axis
type Layer struct {
	Name    string
	Order   int
	sprites []*Sprite
	Scene   *Scene
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
		Scene:   scene,
		sprites: make([]*Sprite, 0, 10), //TODO: make capasity customizable
	}

	l.Scene.layers[name] = &l
	return &l
}

//AddSprite Adds a sprite to the layer. Composite sprites
//can be added by adding the main sprite only
func (l *Layer) AddSprite(sprite *Sprite) {
	sprite.Layer = l
	l.sprites = append(l.sprites, sprite)

}

//Update updates the layer trees
// dt: delta time
func (l *Layer) Update(dt float64) {

	for _, sprite := range l.sprites {
		sprite.Update(dt)
	}
}

//Draw draws the layer tree
func (l *Layer) Draw() {
	//TODO: Add sprite draw functions
	for _, sprite := range l.sprites {
		sprite.Draw()
	}
}

//RemoveSprite a sprite from the layer
func (l *Layer) RemoveSprite(element *Sprite) *Sprite {

	i := -1

	for k, v := range l.sprites {
		if element == v {
			i = k
		}
	}

	if i == -1 {
		return nil
	}

	l.sprites = append(l.sprites[:i], l.sprites[i+1:]...)

	return element

}
