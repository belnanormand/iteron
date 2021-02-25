package iteron

import (
	"github.com/faiface/pixel"
)

//Sprite Stuff that is going to be drawn on screen
type Sprite struct {
	PSprite      *pixel.Sprite
	Behaviours   map[string]interface{ Behaviour }
	Groups       []string
	Layer        *Layer
	SpriteSheet  *pixel.Picture
	Size         Size
	Position     Position
	Frames       []pixel.Rect
	AnchorPoints map[string]Position
	Rotation     float64
	//TODO: Add Velocity to the sprite itslf instead of the behaviours
}

//NewSprite Creates a new sprite and generate frames form the image top left to the bottom right
//
func NewSprite(spritesheet *pixel.Picture, size Size) *Sprite {
	sprite := Sprite{
		SpriteSheet: spritesheet,
		Size:        size,
	}

	sprite.SetPosition(0, 0)
	sprite.Frames = make([]pixel.Rect, 0, 10)
	sprite.Behaviours = make(map[string]interface{ Behaviour })
	sprite.AnchorPoints = make(map[string]Position)
	sprite.Groups = make([]string, 0, 2)

	data := pixel.PictureDataFromPicture(*sprite.SpriteSheet)
	for y := 0.0; y+sprite.Size.Height <= data.Bounds().Max.Y; y += sprite.Size.Height {
		for x := 0.0; x+sprite.Size.Width <= data.Bounds().Max.X; x += sprite.Size.Width {
			sprite.Frames = append(sprite.Frames, pixel.R(
				x,
				y,
				x+sprite.Size.Width,
				y+sprite.Size.Height,
			))
		}
	}

	return &sprite

}

//Update This is where we update the sprite and apply any behaviours attached to the sprite
// dt: delta time
func (sprite *Sprite) Update(dt float64) {
	if sprite.PSprite == nil {
		sprite.PSprite = pixel.NewSprite(nil, pixel.Rect{})
		sprite.PSprite.Set(*sprite.SpriteSheet, sprite.Frames[0])
	}

	for _, behaviour := range sprite.Behaviours {
		behaviour.Apply(dt)
	}

	//TODO: Then apply velocity to position here.

}

//Draw We draw the sprite
func (sprite *Sprite) Draw() {

	if sprite.PSprite == nil {
		return
	} //newly created sprite might not have had their initial update call. Skip if that is the case

	mat := pixel.IM

	mat = mat.Moved(pixel.V(sprite.Position.X, sprite.Position.Y))

	bm, ok := sprite.Behaviours["anchor"].(*BehaviourAnchor)
	if ok {
		mat = mat.Rotated(pixel.V(bm.ParentSprite.Position.X, bm.ParentSprite.Position.Y), -DegreesToRadians(bm.ParentSprite.Rotation))
	}

	mat = mat.Rotated(pixel.V(sprite.Position.X, sprite.Position.Y), -DegreesToRadians(sprite.Rotation))

	sprite.PSprite.Draw(sprite.Layer.Scene.Game.Window, mat)
}

//SetPosition Sets the position of the sprite
func (sprite *Sprite) SetPosition(X float64, Y float64) {
	sprite.Position = Position{X: X, Y: Y}
}

//AddBehaviour Attach behaviours to the sprite
func (sprite *Sprite) AddBehaviour(name string, b interface{ Behaviour }) {
	sprite.Behaviours[name] = b
}

//Destroy Completely removes the sprite from the game
func (sprite *Sprite) Destroy() {
	sprite.Layer.RemoveSprite(sprite)
}
