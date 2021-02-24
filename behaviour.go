package iteron

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

//Behaviour Main interface for behaviours
type Behaviour interface {
	Apply(dt float64)
}

//BehaviourCustom Allows us to crate custom behaviours
type BehaviourCustom struct {
	Behaviour
	Sprite  *Sprite
	Enabled bool
	Update  func(b *BehaviourCustom, dt float64)
}

//Apply calls the attached Update function
func (b *BehaviourCustom) Apply(dt float64) {
	b.Update(b, dt)
}

//Behaviour8Direction Movement behaviour with default WASD
type Behaviour8Direction struct {
	Behaviour
	MaxSpeed     float64 //Maximum speed in p/s
	Acceleration float64 //Acceleration speed in p/s^2
	Deceleration float64 //Deceleration speed in p/s^2
	Sprite       *Sprite
	Enabled      bool
	Velocity     *pixel.Vec
}

//Apply Applies the 8 Derection movement behaviour to a sprite
func (b *Behaviour8Direction) Apply(dt float64) {

	if !b.Enabled {
		return
	}

	w := b.Sprite.Layer.Scene.Game.Window

	if w.Pressed(pixelgl.KeyW) || w.Pressed(pixelgl.KeyS) {
		if w.Pressed(pixelgl.KeyW) {
			b.Velocity.Y += b.Acceleration * dt
		}

		if w.Pressed(pixelgl.KeyS) {
			b.Velocity.Y -= b.Acceleration * dt
		}
	} else {

		if math.Abs(b.Velocity.Y) > 10 {
			if b.Velocity.Y != 0 {
				b.Velocity.Y += (b.Velocity.Y / math.Abs((b.Velocity.Y)) * b.Deceleration * dt)
			}
		} else {
			b.Velocity.Y = 0
		}

	}

	if w.Pressed(pixelgl.KeyA) || w.Pressed(pixelgl.KeyD) {

		if w.Pressed(pixelgl.KeyA) {
			b.Velocity.X -= b.Acceleration * dt
		}

		if w.Pressed(pixelgl.KeyD) {
			b.Velocity.X += b.Acceleration * dt
		}

	} else {

		if math.Abs(b.Velocity.X) > 10 {
			if b.Velocity.X != 0 {
				b.Velocity.X += (b.Velocity.X / math.Abs((b.Velocity.X)) * b.Deceleration * dt)
			}
		} else {
			b.Velocity.X = 0
		}

	}

	if math.Abs(b.Velocity.Y) > b.MaxSpeed {
		b.Velocity.Y = (b.Velocity.Y / math.Abs((b.Velocity.Y))) * b.MaxSpeed
	}

	if math.Abs(b.Velocity.X) > b.MaxSpeed {
		b.Velocity.X = (b.Velocity.X / math.Abs((b.Velocity.X))) * b.MaxSpeed
	}

	b.Sprite.Position.Y += b.Velocity.Y * dt
	b.Sprite.Position.X += b.Velocity.X * dt

}

//BehaviourBoundToLayout Keeps stuff from running offscreen
type BehaviourBoundToLayout struct {
	Behaviour
	Sprite  *Sprite
	Enabled bool
	Margin  float64
}

//Apply Keeps stuff from running offscreen
func (b *BehaviourBoundToLayout) Apply(dt float64) {
	w := b.Sprite.Layer.Scene.Game.Window

	if b.Sprite.Position.Y > (w.Bounds().H() - b.Margin) {
		b.Sprite.Position.Y = (w.Bounds().H() - b.Margin)
	}

	if b.Sprite.Position.X > (w.Bounds().W() - b.Margin) {
		b.Sprite.Position.X = (w.Bounds().W() - b.Margin)
	}

	if b.Sprite.Position.Y < b.Margin {
		b.Sprite.Position.Y = b.Margin
	}

	if b.Sprite.Position.X < b.Margin {
		b.Sprite.Position.X = b.Margin
	}
}

//BehaviourAnimation VERY basic animation behaviour. Loops the frames from back to front
type BehaviourAnimation struct {
	Behaviour
	Sprite           *Sprite
	accumelator      float64
	FrameTime        float64
	currentFrame     int
	Enabled          bool
	AnimationSet     map[string][]pixel.Rect
	CurrentAnimation string
}

//Apply VERY basic animation behaviour. Loops the frames from back to front
func (b *BehaviourAnimation) Apply(dt float64) {
	b.accumelator += dt
	if b.accumelator > b.FrameTime {
		b.accumelator -= b.FrameTime
		b.currentFrame++

		if len(b.AnimationSet) > 0 && len(b.CurrentAnimation) > 0 {
			if b.currentFrame > len(b.AnimationSet[b.CurrentAnimation])-1 {
				b.currentFrame = 0
			}
			b.Sprite.PSprite.Set(*b.Sprite.SpriteSheet, b.AnimationSet[b.CurrentAnimation][b.currentFrame])
		} else {
			if b.currentFrame >= len(b.Sprite.Frames) {
				b.currentFrame = 0
			}
			b.Sprite.PSprite.Set(*b.Sprite.SpriteSheet, b.Sprite.Frames[b.currentFrame])
		}

	}
}

//BehaviourAnchor Anchors one sprite to another
type BehaviourAnchor struct {
	Behaviour
	Sprite       *Sprite
	Enabled      bool
	ParentSprite *Sprite
	AnchorName   string
}

//Apply Anchors a sprite to another
func (b *BehaviourAnchor) Apply(dt float64) {
	b.Sprite.Position.X = b.ParentSprite.Position.X + b.ParentSprite.AnchorPoints[b.AnchorName].X
	b.Sprite.Position.Y = b.ParentSprite.Position.Y + b.ParentSprite.AnchorPoints[b.AnchorName].Y
}

//BehaviourBullet Moves sprite in a set direction
type BehaviourBullet struct {
	Behaviour
	Speed     float64
	Sprite    *Sprite
	Enabled   bool
	Velocity  *pixel.Vec
	Direction float64 //(0 - 360) (360 - 720)
}

//Apply Moves sprite in a set direction
func (b *BehaviourBullet) Apply(dt float64) {

	b.Velocity.Y = b.Speed * math.Cos(DegreesToRadians(b.Direction))
	b.Velocity.X = b.Speed * math.Sin(DegreesToRadians(b.Direction))

	b.Sprite.Position.Y += b.Velocity.Y * dt
	b.Sprite.Position.X += b.Velocity.X * dt

}
