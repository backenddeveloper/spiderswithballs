package framework

//resolveCollisionWith finds a point of collision between two sprites
// the 's sprite' has its position adjusted, the 'o other reference sprite' does not
func (s *Sprite) resolveCollisionWith (o *Sprite) {

	// Here we take a mean average of the two coefficients of restitution
	// https://en.wikipedia.org/wiki/Coefficient_of_restitution
	// This is roughly in line with how real physics works
	e := (s.e + o.e) / 2

	// the collison ratio is the fraction of the distance between the two surfaces at which the collision occurs

	// S |-----X--------------| O
	collidesX, ratioX, positionX := collisionLocation(s.PositionX, s.Width, s.SpeedX, o.PositionX, o.Width, o.SpeedX)

	// S |-----Y--------------| O
	collidesY, ratioY, positionY := collisionLocation(s.PositionY, s.Width, s.SpeedY, o.PositionY, o.Width, o.SpeedY)

	// if the two sprites collide in both the x and y planes
	if collidesX && collidesY {

		// here we 
	}
}

// Find any collisions in the X direction
func collisionLocation (sPosition, sWidth, sSpeed, oPosition, oWidth, oSpeed float64) collides bool, collisionRatio, collisionPosition, collisionCorrection float64 {

	// if S is to the left of O
	if sPosition + sWidth < oPosition {

		// if the relative speed between them is greater than the distance
		if (sSpeed - oSpeed) > (oPosition - (sPosition + sWidth)) {

			// we note that a collision takes place
			collides = true

			// this is the fraction of the distance between the two surfaces at which the collision occurs
			// S |-----X--------------| O
			collisionRatio = ((oPosition - (sPosition + sWidth)) / (sSpeed - oSpeed)) + (sPosition + sWidth)
			collisionPosition = (1 - collisionRatio)*(sPosition + sWidth) + oPosition
			collisionCorrection = collisionPosition - oPosition
		}

	// if O is to the left of S
	} else if oPosition + oWidth < sPosition {

		// if the relative speed between them is greater than the distance
		if (oSpeed - sSpeed) > (sPosition - (oPosition + oWidth)) {

			// we note that a collision takes place
			collides = true

			// this is the fraction of the distance between the two surfaces at which the collision occurs
			// O |-----X--------------| S
			collisionRatio = ((sPosition - (oPosition + oWidth)) / (oSpeed - sSpeed)) + (oPosition + oWidth)
			collisionPosition = (1 - collisionRatio)*(oPosition + oWidth) + sPosition
			collisionCorrection = collisionPosition - sPosition
		}

	// else there is initial overlap, this indicates a 'normal angle' collision 
	} else {

		// Sprites really shouldn't be moving this fast!
		collides = true
	}
}
