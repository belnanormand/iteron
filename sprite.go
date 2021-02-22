package iteron

import (
	"github.com/faiface/pixel"
)

//Sprite Stuff that is going to be drawn on screen
type Sprite struct {
	pSprite     *pixel.Sprite
	behaviours  []interface{ Behaviour }
	sprites     []*Sprite
	layer       *Layer
	spriteSheet *pixel.Picture
	size        Size
	position    Position
	Frames      []pixel.Rect
	//TODO: Add Velocity to the sprite itslf instead of the behaviours
}

//NewSprite Creates a new sprite and generate frames form the image top left to the bottom right
//
func NewSprite(spritesheet *pixel.Picture, size Size) *Sprite {
	sprite := Sprite{
		spriteSheet: spritesheet,
		size:        size,
	}

	sprite.Frames = make([]pixel.Rect, 0, 10)
	sprite.behaviours = make([]interface{ Behaviour }, 0, 10)

	data := pixel.PictureDataFromPicture(*sprite.spriteSheet)
	for y := 0.0; y+sprite.size.Height <= data.Bounds().Max.Y; y += sprite.size.Height {
		for x := 0.0; x+sprite.size.Width <= data.Bounds().Max.X; x += sprite.size.Width {
			sprite.Frames = append(sprite.Frames, pixel.R(
				x,
				y,
				x+sprite.size.Width,
				y+sprite.size.Height,
			))
		}
	}

	return &sprite

}

//Update This is where we update the sprite and apply any behaviours attached to the sprite
// dt: delta time
func (sprite *Sprite) Update(dt float64) {
	if sprite.pSprite == nil {
		sprite.pSprite = pixel.NewSprite(nil, pixel.Rect{})
		sprite.pSprite.Set(*sprite.spriteSheet, sprite.Frames[0])
	}

	for _, behaviour := range sprite.behaviours {
		behaviour.Apply(dt)
	}

	//TODO: Then apply velocity to position here.

}

//Draw We draw the sprite
func (sprite *Sprite) Draw() {
	sprite.pSprite.Draw(sprite.layer.scene.game.window,
		pixel.IM.Moved(pixel.V(sprite.position.X, sprite.position.Y)))
}

//SetPosition Sets the position of the sprite
func (sprite *Sprite) SetPosition(X float64, Y float64) {
	sprite.position = Position{X: X, Y: Y}
}

//AddBehaviour Attach behaviours to the sprite
func (sprite *Sprite) AddBehaviour(b interface{ Behaviour }) {
	sprite.behaviours = append(sprite.behaviours, b)
}
