package actor

import (
	"github.com/bjatkin/flappy_boot/internal/game"
	"github.com/bjatkin/flappy_boot/internal/math"
)

// Player is a struct representing a player
type Player struct {
	Sprite *game.Sprite
	dy     math.Fix8
	maxDy  math.Fix8

	dead    bool
	started bool
}

// NewPlayer creates a new player struct
func NewPlayer(x, y math.Fix8, sprite *game.Sprite) *Player {
	p := &Player{
		Sprite: sprite,
		maxDy:  math.FixOne * 8,
	}

	p.Sprite.TileIndex = 16
	p.Sprite.SetAnimation(glideAni)
	p.Sprite.X = x
	p.Sprite.Y = y
	return p
}

// Start indicates the the game has started and the player should start applying physics
func (p *Player) Start() {
	p.started = true
}

// Dead sets the player state to dead
func (p *Player) Dead() {
	p.dead = true
}

// Reset resets all the players properties to be the same as they were on creation. It also move the sprite to the specified location
func (p *Player) Reset(x, y math.Fix8) {
	p.dy = 0
	p.started = false
	p.dead = false
	p.Sprite.X = x
	p.Sprite.Y = y
	p.Sprite.HFlip = false
	p.Sprite.TileIndex = 16
	p.Sprite.SetAnimation(glideAni)
}

// Rect returns the hitbox of the player as a math.Rect
func (p *Player) Rect() math.Rect {
	return math.Rect{
		X1: p.Sprite.X.Int() + 12,
		Y1: p.Sprite.Y.Int() + 2,
		X2: p.Sprite.X.Int() + 22,
		Y2: p.Sprite.Y.Int() + 12,
	}
}

// Show whos the player sprite
func (p *Player) Show() error {
	err := p.Sprite.Add()
	if err != nil {
		return err
	}

	return nil
}

// Hide hides the player sprite
func (p *Player) Hide() {
	p.Sprite.Remove()
}

var (
	jumpAni = []game.Frame{
		{Index: 16, Len: 3},
		{Index: 32, Len: 4},
		{Index: 0, Len: 7},
		{Index: 8, Len: 8},
		{Index: 24, Len: 2},

		{Index: 16, Len: 40},
		{Index: 24, Len: 40, Offset: math.V2{X: 0, Y: math.FixOne}},

		{Index: 16, Len: 40},
		{Index: 24, Len: 40, Offset: math.V2{X: 0, Y: math.FixOne}},

		{Index: 16, Len: 40},
		{Index: 24, Len: 40, Offset: math.V2{X: 0, Y: math.FixOne}},
	}

	glideAni = []game.Frame{
		{Index: 16, Len: 40},
		{Index: 24, Len: 40, Offset: math.V2{X: 0, Y: math.FixOne}},
	}
)

// Update updates the players physics and interal properites
func (p *Player) Update(gravity, jump math.Fix8) {
	p.Sprite.Update()

	if !p.started {
		// don't update physics if the game has not started yet
		return
	}

	p.dy += gravity
	if p.dy > p.maxDy {
		p.dy = p.maxDy
	}

	if jump != 0 {
		p.Sprite.SetAnimation(jumpAni)
		p.dy = jump
	}

	if p.dead {
		p.Sprite.StopAnimation()
		p.Sprite.HFlip = true
		p.Sprite.TileIndex = 0
	}

	p.Sprite.Y += p.dy
	if p.Sprite.Y > math.FixOne*200 {
		p.Sprite.Y = math.FixOne * 200
	}

	if p.Sprite.Y < -math.FixOne*16 {
		p.Sprite.Y = -math.FixOne * 16
	}
}
