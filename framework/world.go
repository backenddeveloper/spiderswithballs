package framework

//World is a collection of sprites and the Physics model for sprites colliding
type World struct {
	SpriteList []*Sprite
}

//NewWorld creates a new World object containing a maximum of 8192 Sprites
func NewWorld() *World {
	return &World{
		make([]*Sprite, 0, 8192),
	}
}

//Update performs one 'clock' of all the world's Sprites' positions
// it is up to the consumer to determine how frequently these happen
// Update adjusts the particle's speed for gravity and bounce then updates the sprite
func (w *World) Update() {

	// First we deep copy of the sprite array to freeze it as a reference
	copySpriteList := make([]*Sprite, len(w.SpriteList))
	copy(copySpriteList, w.SpriteList)

	// Then we update the speeds and relative postions of the sprites
	for _, s := range w.SpriteList {

		// accelerate due to gravity. Note that the framework deals in positions as
		// percentages, with 0,0 being the top left corner and 100,100 being the bottom right
		// Therefore gravity is added to vertical speed with each update
		s.SpeedY += s.GravitationalConstant

		// if the sprite attempts to go below impact with the floor it bounces
		// we also check it is moving in a downwards direction to prevent a sprite
		// rendered below the floor from vibrating.
		if s.PositionY >= (100-s.Height) && s.SpeedY > 0.001 {

			// https://en.wikipedia.org/wiki/Coefficient_of_restitution
			s.SetYSpeed(-s.SpeedY * s.CoefficientOfRestitution)
		}

		// Here we update the speed of the sprite based on inter-sprite collisions
		// 'o' for 'other sprite'
		for _, o := range copySpriteList {
			if s != o {

				// resolve collisions
				s.resolveCollisionWith(o)
			}
		}

		// This updates the Sprite's position
		s.PositionX += s.SpeedX
		s.PositionY += s.SpeedY
	}
}

//AddSprite adds a sprite to the world
func (w *World) AddSprite(s *Sprite) int {
	w.SpriteList = append(w.SpriteList, s)
	return len(w.SpriteList)
}
